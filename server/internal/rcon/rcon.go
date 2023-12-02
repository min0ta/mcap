package rcon

import "github.com/gorcon/rcon"

type RconClient struct {
	conn *rcon.Conn
}

func New() *RconClient {
	return &RconClient{}
}

func (r *RconClient) Dial() error {
	var err error
	r.conn, err = rcon.Dial("nomeg.ru:25569", "123test123")
	if err != nil {
		return err
	}
	return nil
}

func (r *RconClient) Exec(command string) (response string, e error) {
	return r.conn.Execute(command)
}
