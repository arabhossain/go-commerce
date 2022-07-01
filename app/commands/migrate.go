package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/joho/godotenv"
	"go-commerce/app/configs"
	"go-commerce/app/utils"
	"log"
)

func main() {
	godotenv.Load()

	app, err := configs.GetInstance()
	if err != nil {
		utils.Error.Fatal(err.Error())
	}
	direction := app.Cfg.GetMigration()
	if direction != "down" && direction != "up" {
		log.Println("-migrate accepts [up, down] values only")
		return
	}

	//db, _ := sql.Open("mysql", "user:password@tcp(host:port)/dbname?multiStatements=true")
	db, _ := sql.Open("mysql", app.Cfg.GetDBConnStr())
	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		"file://app/database/migrations",
		"mysql",
		driver,
	)

	if direction == "up" {
		if err := m.Up(); err != nil {
			log.Printf("failed migrate up: %s", err)
			return
		}
	}

	if direction == "down" {
		if err := m.Down(); err != nil {
			log.Printf("failed migrate down: %s", err)
			return
		}
	}

	log.Printf("Migration has been completed. Direction: %s", direction)
}
