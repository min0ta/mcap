package main

import (
	"mcap/internal/auth"
	"mcap/internal/config"
)

func main() {
	cfg := config.New()
	authh := auth.New(cfg)
	authh.Test()
}
