package api

import (
	"io"
	"mcap/internal/config"
	"net/http"
)

type ApiServer struct {
	cfg *config.Config
}

func New(config *config.Config) *ApiServer {
	return &ApiServer{
		cfg: config,
	}
}

func (s *ApiServer) Start() error {
	s.configureRouter()
	if err := http.ListenAndServe(s.cfg.ServerPort, nil); err != nil {
		return err
	}
	return nil
}

func (s *ApiServer) configureRouter() {
	http.HandleFunc("/test", s.testRoute)
}

func (s *ApiServer) testRoute(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello world!")
}
