package api

import (
	"mcap/internal/auth"
	"mcap/internal/config"
	"net/http"
)

type ApiServer struct {
	cfg           *config.Config
	authorization *auth.Authoriztaion
}

func New(config *config.Config) *ApiServer {
	return &ApiServer{
		cfg: config,
	}
}

func (s *ApiServer) Start() error {
	s.authorization = auth.New(s.cfg)
	s.configureRouter()
	return http.ListenAndServe(s.cfg.SERVER_PORT, nil)
}

func (s *ApiServer) configureRouter() {
	http.HandleFunc("/login", s.authorization.Authorize)
	http.HandleFunc("/test", s.authorization.TestIfAuth)
}
