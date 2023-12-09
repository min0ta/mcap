package api

import (
	"mcap/internal/auth"
	"mcap/internal/config"
	"mcap/internal/errors"
	"mcap/internal/log"
	mcservermanager "mcap/internal/mcServerManager"
	"mcap/internal/utils"
	"net/http"
)

type ApiServer struct {
	cfg           *config.Config
	authorization *auth.Authoriztaion
	logger        *log.Logger
	mcServers     []*mcservermanager.MinecraftServer
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
	serverConfigs := mcservermanager.ParseConfigs("./config/servers.json")

	for i := range serverConfigs {
		mcServer := mcservermanager.New(serverConfigs[i])
		s.mcServers = append(s.mcServers, mcServer)
	}

	s.logger.WriteFormat("server started at port %s", s.cfg.SERVER_PORT)
	return http.ListenAndServe(s.cfg.SERVER_PORT, nil)
}

func (s *ApiServer) configureRouter() {
	http.HandleFunc("/login", s.authorization.Authorize)
	http.HandleFunc("/servers", s.showServers)
	http.HandleFunc("/start", s.startServer)
	http.HandleFunc("/server", s.getServer)
	if s.cfg.TEST_ROUTE {
		http.HandleFunc("/test", s.authorization.TestIfAuth)
		http.HandleFunc("/rcon", s.execRcon)
		http.HandleFunc("/stop", s.stopServer)
	}
}

type execQuery struct {
	Server  string `json:"server"`
	Command string `json:"command"`
}

func (s *ApiServer) execRcon(w http.ResponseWriter, r *http.Request) {
	role := s.authorization.AuthCheck(r)
	if role != auth.RoleAdmin {
		errors.HttpError(w, errors.ErrorUnauthorized, 401)
		return
	}
	q := &execQuery{}

	err := utils.ReadJson(r, q)
	if err != nil {
		errors.HttpError(w, errors.ErrorInvalidQuery, 400)
		return
	}

	serverIndex := utils.Find(s.mcServers, func(ms *mcservermanager.MinecraftServer) bool {
		return ms.Config.Name == q.Server
	})
	if serverIndex == -1 {
		errors.HttpError(w, errors.ErrorInvalidQuery, 400)
		return
	}

	res, err := s.mcServers[serverIndex].Rcon.Exec(q.Command)
	if err != nil {
		s.logger.WriteFormat("cannot execute rcon with error %v", err)
		errors.HttpError(w, errors.ErrorCannotAccessRcon, 500)
		return
	}

	utils.WriteResult(w, utils.Response{
		"res": res,
	}, 200)
}

type serverStartQuery struct {
	Server string `json:"server"`
}

func (s *ApiServer) startServer(w http.ResponseWriter, r *http.Request) {
	role := s.authorization.AuthCheck(r)
	if role == auth.RoleGuest {
		errors.HttpError(w, errors.ErrorUnauthorized, 401)
		return
	}

	q := &serverStartQuery{}
	err := utils.ReadJson(r, q)
	if err != nil {
		errors.HttpError(w, errors.ErrorInvalidQuery, 400)
		return
	}
	index := utils.Find(s.mcServers, func(ms *mcservermanager.MinecraftServer) bool {
		return ms.Config.Name == q.Server
	})
	if index == -1 {
		errors.HttpError(w, errors.ErrorInvalidQuery, 400)
		return
	}
	err = s.mcServers[index].Start()
	if err != nil {
		s.logger.WriteFormat("cannot start server with error %v", err)
		errors.HttpError(w, errors.ErrorCannotStartMcServer, 500)
		return
	}
	utils.WriteResult(w, utils.Response{
		"success": true,
	}, 200)
}

func (s *ApiServer) stopServer(w http.ResponseWriter, r *http.Request) {
	role := s.authorization.AuthCheck(r)
	if role == auth.RoleGuest {
		errors.HttpError(w, errors.ErrorUnauthorized, 401)
		return
	}
	q := serverStartQuery{}
	err := utils.ReadJson(r, &q)
	if err != nil {
		errors.HttpError(w, errors.ErrorInvalidQuery, 400)
		return
	}
	index := utils.Find(s.mcServers, func(mc *mcservermanager.MinecraftServer) bool {
		return mc.Config.Name == q.Server
	})
	if index == -1 {
		errors.HttpError(w, errors.ErrorInvalidQuery, 400)
		return
	}
	err = s.mcServers[index].Stop()
	if err != nil {
		s.logger.WriteFormat("cannot start server with error %v", err)
		errors.HttpError(w, errors.ErrorCannotStartMcServer, 500)
		return
	}
	utils.WriteResult(w, utils.Response{
		"success": true,
	}, 200)
}

type serversListResponse struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Port    string `json:"port"`
	Online  bool   `json:"online"`
}

func (s *ApiServer) showServers(w http.ResponseWriter, r *http.Request) {
	role := s.authorization.AuthCheck(r)
	if role == auth.RoleGuest {
		errors.HttpError(w, errors.ErrorUnauthorized, 401)
		return
	}
	var displayServers []serversListResponse
	for i := range s.mcServers {
		server := s.mcServers[i].Config
		displayServers = append(displayServers, serversListResponse{
			Name:    server.Name,
			Address: server.Address,
			Port:    server.Port,
			Online:  s.mcServers[i].IsOnline,
		})
	}
	utils.WriteResult(w, utils.Response{
		"list": displayServers,
	}, 200)
}

func (s *ApiServer) getServer(w http.ResponseWriter, r *http.Request) {
	role := s.authorization.AuthCheck(r)
	if role == auth.RoleGuest {
		errors.HttpError(w, errors.ErrorUnauthorized, 401)
		return
	}
	q := &serverStartQuery{}
	err := utils.ReadJson(r, q)
	if err != nil {
		errors.HttpError(w, errors.ErrorInvalidQuery, 400)
		return
	}
	index := utils.Find(s.mcServers, func(ms *mcservermanager.MinecraftServer) bool {
		return q.Server == ms.Config.Name
	})
	if index == -1 {
		errors.HttpError(w, errors.ErrorInvalidQuery, 400)
		return
	}
	utils.WriteResult(w, utils.Response{
		"online": s.mcServers[index].IsOnline,
		"name":   q.Server,
	}, 200)
}
