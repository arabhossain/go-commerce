package configs

import "go-commerce/app/utils"

// Application holds commonly used app wide data, for ease of DI
type Application struct {
	DB  *utils.DB
	Cfg *Config
}

// GetInstance Get captures env vars, establishes DB connection and keeps/returns
// reference to both
func GetInstance() (*Application, error)  {
	cfg := Get()

	db, err := utils.Get(cfg.GetDBConnStr())
	if err != nil {
		return nil, err
	}

	return &Application{
		DB:  db,
		Cfg: cfg,
	}, nil
}

