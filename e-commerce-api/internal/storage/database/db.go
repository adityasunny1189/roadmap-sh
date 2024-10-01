package database

import (
	"database/sql"

	"github.com/adityasunny1189/roadmap-sh/e-commerce-api/internal/common/config"
	"github.com/go-sql-driver/mysql"
)

func Load(config *config.Config) *sql.DB {
	cfg := mysql.Config{
		User:   config.DBDetails.User,
		Passwd: config.DBDetails.Password,
		Net:    config.DBDetails.Net,
		Addr:   config.DBDetails.Address,
		DBName: config.DBDetails.DBName,
	}
	db, err := sql.Open(config.DBDetails.DB, cfg.FormatDSN())
	if err != nil {
		panic(err)
	}

	return db
}
