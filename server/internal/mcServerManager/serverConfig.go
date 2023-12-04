package mcservermanager

import (
	"encoding/json"
	"log"
	"mcap/internal/utils"
)

type ServerConfig struct {
	Name         string `json:"name"`
	RunCommand   string `json:"runCommand"`
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
		pServers = append(pServers, &servers[k])
	}
	return pServers
}
