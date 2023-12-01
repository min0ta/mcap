package api

import (
	"mcap/internal/auth"
	"mcap/internal/config"
	"mcap/internal/log"
	"net/http"
)

type ApiServer struct {
	cfg           *config.Config
	authorization *auth.Authoriztaion
	logger        *log.Logger
}

func New(config *config.Config) *ApiServer {
	return &ApiServer{
		cfg:    config,
		logger: log.New(log.BothMode, "server.log"),
	}
}

func (s *ApiServer) Start() error {
	s.authorization = auth.New(s.cfg, s.logger)
	s.configureRouter()
	s.logger.WriteFormat("server started at port %s", s.cfg.SERVER_PORT)
	return http.ListenAndServe(s.cfg.SERVER_PORT, nil)
}

func (s *ApiServer) configureRouter() {
	http.HandleFunc("/login", s.authorization.Authorize)
	http.HandleFunc("/test", s.authorization.TestIfAuth)
}
