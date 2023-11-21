package api

import "mcap/internal/config"

type ApiServer struct {
	cfg *config.Config
}

func New(config *config.Config) *ApiServer {
	return &ApiServer{
		cfg: config,
	}
}
