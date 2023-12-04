package config

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

type Config struct {
	SERVER_PORT     string
	JWT_SIGNING_KEY string
	JWT_EXPIRES     uint
	PATH_TO_JSON_DB string
	LOG_MODE        string
	LOG_FILE        string
	TEST_ROUTE      bool
	RCON_ADDRESS    string
	RCON_PASSWORD   string
}

func New() *Config {
	return &Config{
		SERVER_PORT:     ":8080",
		JWT_SIGNING_KEY: "jwt signing key! should be random string of characters!",
		JWT_EXPIRES:     86400,
		PATH_TO_JSON_DB: "./config/db.json",
		LOG_MODE:        "both",
		LOG_FILE:        "server.log",
		TEST_ROUTE:      true,
		RCON_ADDRESS:    "put your rcon address here",
		RCON_PASSWORD:   "put your rcon password here",
	}
}

func (cfg *Config) ReadJsonConfig(path string) {
	_json, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll("config", os.ModePerm)
		if err != nil {
			log.Fatal("cannot create config dir with error", err)
		}
		file, err := os.Create("config/config.json")
		if err != nil {
			log.Fatal("cannot create config file", err)
		}
		defer file.Close()
		__json, err := json.Marshal(cfg)
		if err != nil {
			log.Fatal("cannot marshal cfg ", err)
		}
		file.Write(__json)
		_json = __json
	} else if err != nil {
		log.Fatal("cannot open cfg file ", err)
	}
	if err := json.Unmarshal(_json, cfg); err != nil {
		log.Fatal("unable to parse json config with error: ", err)
	}
}
