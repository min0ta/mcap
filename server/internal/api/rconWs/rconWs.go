package rconws

import (
	"mcap/internal/errors"
	mcservermanager "mcap/internal/mcServerManager"
	"net/http"

	"github.com/gorilla/websocket"
)

type Message struct {
	text  string
	yours bool
}

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

func HandleConns(server *mcservermanager.MinecraftServer, w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		errors.HttpError(w, errors.ErrorCannotUpgrade, 400)
		return
	}
	defer ws.Close()
	for {
		msg := server.ReadLogs()
		err = ws.WriteJSON(Message{
			text:  msg,
			yours: false,
		})
		if err != nil {
			errors.HttpError(w, errors.ErrorUnknow, 400)
			return
		}
	}
}
