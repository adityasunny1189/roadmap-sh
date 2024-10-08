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

func NewConfig() *Config {
	return &Config{
		DBDetails: MysqlDBDetails{
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Net:      "tcp",
			Address:  "host.docker.internal:3306", //"127.0.0.1:3306",
			DBName:   "my_commerce",
			DB:       "mysql",
		},
		RedisUrlPath: "",
	}
}
