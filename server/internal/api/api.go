package api

import (
	"fmt"
	"io"
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
	fmt.Println("server running at port ", s.cfg.SERVER_PORT)
	s.authorization = auth.New(s.cfg, s.logger)
	s.configureRouter()
	s.logger.WriteString("server started")
	return http.ListenAndServe(s.cfg.SERVER_PORT, nil)
}

func (s *ApiServer) configureRouter() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "listenin'...")
		w.WriteHeader(404)
		fmt.Println("access ", r.URL)
	})
	http.HandleFunc("/login", s.authorization.Authorize)
	http.HandleFunc("/test", s.authorization.TestIfAuth)
}
