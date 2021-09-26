package main

import (
	"log"

	"github.com/stalin-777/accounting-of-jobs/config"
	"github.com/stalin-777/accounting-of-jobs/logger"
	"github.com/stalin-777/accounting-of-jobs/postgres"
	"github.com/stalin-777/accounting-of-jobs/server"
)

func main() {

	cfg, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	err = logger.InitZapLogger(
		cfg.Logger.Path,
		cfg.Logger.FileName,
		cfg.Logger.MaxSize,
		cfg.Logger.MaxBackups,
		cfg.Logger.MaxAge,
	)
	if err != nil {
		log.Fatal(err)
	}

	connPool, err := postgres.NewPgxConnPool(cfg.DB.Host, cfg.DB.Port, cfg.DB.Username, cfg.DB.Password, cfg.DB.Dbname)
	if err != nil {
		logger.Fatal(err)
	}

	err = postgres.Migrate(cfg.DB.Host, cfg.DB.Port, cfg.DB.Username, cfg.DB.Password, cfg.DB.Dbname)
	if err != nil {
		logger.Fatal(err)
	}

	server.Run(cfg.WebPort, connPool)
}
