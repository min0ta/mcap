package mcservermanager

import (
	"encoding/json"
	"log"
	"mcap/internal/utils"
	"strings"
)

type ServerConfig struct {
	Name         string `json:"name"`
	RunCommand   string `json:"runCommand"`
	Args         []string
	Address      string `json:"address"`
	Port         string `json:"port"`
	RconAddress  string `json:"rconAddress"`
	RconPassword string `json:"rconPwd"`
}

func ParseConfigs(path string) []*ServerConfig {
	_json := utils.RequireFile(path)
	servers := []ServerConfig{}
	err := json.Unmarshal(_json, &servers)
	if err != nil {
		log.Fatal("cannot unmarshall server config with error ", err)
	}
	pServers := []*ServerConfig{}
	for k := range servers {
		splitted := strings.Split(servers[k].RunCommand, " ")
		servers[k].RunCommand = splitted[0]
		servers[k].Args = splitted[1:]
		pServers = append(pServers, &servers[k])
	}
	return pServers
}
