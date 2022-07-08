package main

import (
	"github.com/joho/godotenv"
	"go-commerce/app/configs"
	"go-commerce/app/utils"
	"net"
)

func main() {
	if err := godotenv.Load(); err != nil {
		utils.Info.Println("failed to load env vars")d
	}

	app, err := configs.GetInstance()
	if err != nil {
		utils.Error.Fatal(err.Error())
	}

	srv := utils.GetServer().
		WithAddr(app.Cfg.GetAPIPort()).
		WithRouter(configs.SetupRouter()).
		WithErrLogger(utils.Error)


	utils.Info.Printf("starting server at http://%s%s", GetLocalIP(),app.Cfg.GetAPIPort())
	if err := srv.Start(); err != nil {
		utils.Error.Fatal(err.Error())
	}

	utils.ExitHandler(func() {
		if err := srv.Close(); err != nil {
			utils.Error.Println(err.Error())
		}

		if err := app.DB.Close(); err != nil {
			utils.Error.Println(err.Error())
		}
	})

}

func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return "localhost"
}

