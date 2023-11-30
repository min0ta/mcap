package main

import (
	"flag"
	"fmt"
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
	fmt.Println("test")
	flag.Parse()

	cfg := config.New()
	cfg.ReadJsonConfig(configPath)

	server := api.New(cfg)
	err := server.Start()
	if err != nil {
		log.Fatal(err)
	}
}
