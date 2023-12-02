package rcon

import (
	"mcap/internal/config"

	"github.com/gorcon/rcon"
)

type RconClient struct {
	conn *rcon.Conn
	cfg  *config.Config
}

func New(c *config.Config) *RconClient {
	return &RconClient{
		cfg: c,
	}
}

func (r *RconClient) Dial() error {
	var err error
	r.conn, err = rcon.Dial(r.cfg.RCON_ADDRESS, r.cfg.RCON_PASSWORD)
	if err != nil {
		return err
	}
	return nil
}

func (r *RconClient) Exec(command string) (response string, e error) {
	return r.conn.Execute(command)
}
