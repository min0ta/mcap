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
		logger: log.New(config.LOG_MODE, config.LOG_FILE),
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
	if s.cfg.TEST_ROUTE {
		http.HandleFunc("/test", s.authorization.TestIfAuth)
	}
}
