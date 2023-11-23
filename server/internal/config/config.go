package config

import (
	"encoding/json"
	"log"
	"mcap/internal/utils"
)

type Config struct {
	SERVER_PORT     string
	JWT_SIGNING_KEY string
	JWT_EXPIRES     uint
	PATH_TO_JSON_DB string
}

func New() *Config {
	return &Config{
		SERVER_PORT:     ":8080",
		JWT_SIGNING_KEY: "jwt signing key! should be random string of characters!",
		JWT_EXPIRES:     86400,
		PATH_TO_JSON_DB: "./config/db.json",
	}
}

func (cfg *Config) ReadJsonConfig(path string) {
	_json := utils.RequireFile(path)
	if err := json.Unmarshal(_json, cfg); err != nil {
		log.Fatal("unable to parse json config with error: ", err)
	}
}
