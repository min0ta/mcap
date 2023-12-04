package rcon

import (
	"github.com/gorcon/rcon"
)

type RconClient struct {
	conn    *rcon.Conn
	address string
	pwd     string
}

func New(address string, pwd string) *RconClient {
	return &RconClient{
		address: address,
		pwd:     pwd,
	}
}

func (r *RconClient) Dial() error {
	var err error
	r.conn, err = rcon.Dial(r.address, r.pwd)
	if err != nil {
		return err
	}
	return nil
}

func (r *RconClient) Exec(command string) (response string, e error) {
	return r.conn.Execute(command)
}
