package config

import "os"

type Config struct {
	DBDetails    MysqlDBDetails
	RedisUrlPath string
}

type MysqlDBDetails struct {
	User     string
	Password string
	Net      string
	Address  string
	DBName   string
	DB       string
}

func LoadConfig() *Config {
	return &Config{
		DBDetails: MysqlDBDetails{
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Net:      "tcp",
			Address:  "127.0.0.1:3306", //"host.docker.internal:3306", //
			DBName:   "auth",
			DB:       "mysql",
		},
		RedisUrlPath: "",
	}
}
