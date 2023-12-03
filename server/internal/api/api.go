package api

import (
	"mcap/internal/auth"
	"mcap/internal/config"
	"mcap/internal/errors"
	"mcap/internal/log"
	"mcap/internal/rcon"
	"mcap/internal/utils"
	"net/http"
)

type ApiServer struct {
	cfg           *config.Config
	authorization *auth.Authoriztaion
	logger        *log.Logger
	rcon          *rcon.RconClient
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
	_rcon := rcon.New(s.cfg)
	err := _rcon.Dial()
	if err != nil {
		s.logger.WriteFormat("cannot connect to rcon with error %v", err)
	}
	s.logger.WriteFormat("server started at port %s", s.cfg.SERVER_PORT)
	return http.ListenAndServe(s.cfg.SERVER_PORT, nil)
}

func (s *ApiServer) configureRouter() {
	http.HandleFunc("/login", s.authorization.Authorize)
	if s.cfg.TEST_ROUTE {
		http.HandleFunc("/test", s.authorization.TestIfAuth)
	}
	http.HandleFunc("/rcon", s.execRcon)
}

type execQuery struct {
	Command string `json:"command"`
}

func (s *ApiServer) execRcon(w http.ResponseWriter, r *http.Request) {
	role := s.authorization.AuthCheck(r)
	if role != auth.RoleAdmin {
		errors.HttpError(w, errors.ErrorUnauthorized, 401)
		s.logger.WriteFormat("RCON ACCESS DENY! REQUEST FROM IP: %s", r.RemoteAddr)
		return
	}
	q := &execQuery{}
	err := utils.ReadJson(r, q)
	if err != nil {
		errors.HttpError(w, errors.ErrorInvalidQuery, 400)
		return
	}
	res, err := s.rcon.Exec(q.Command)
	if err != nil {
		errors.HttpError(w, errors.ErrorCannotAccessRcon, 500)
		return
	}
	utils.WriteResult(w, utils.Response{
		"res": res,
	}, 200)
}
