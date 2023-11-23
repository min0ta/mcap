package main

import (
	"flag"
	"log"
	"mcap/internal/api"
	"mcap/internal/config"
)

func init() {
	flag.StringVar(&configPath, "config", "config/config.json", "path to json config file")
	flag.StringVar(&configPath, "c", "config/config.json", "path to json config file")
}

var (
	configPath string
)

func main() {
	flag.Parse()

	cfg := config.New()
	cfg.ReadJsonConfig(configPath)

	server := api.New(cfg)
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
