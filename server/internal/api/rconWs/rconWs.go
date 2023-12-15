package rconws

import (
	"fmt"
	"mcap/internal/errors"
	mcservermanager "mcap/internal/mcServerManager"
	"mcap/internal/utils"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type Message struct {
	Text  string
	Yours bool
}
type InitQuery struct {
	Server string `json:"server"`
}

type Req struct {
	Type string `json:"type"`
}

type Res struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

const (
	TypeKeepAlive = "keep"
	TypeLogs      = "logs"
)

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

func GetLogs(servers []*mcservermanager.MinecraftServer, w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		errors.HttpError(w, errors.ErrorCannotUpgrade, 400)
		return
	}
	defer ws.Close()
	q := InitQuery{}
	err = ws.ReadJSON(&q)
	if err != nil {
		ws.WriteJSON(utils.Response{
			"err": errors.ErrorInvalidQuery,
		})
		return
	}
	index := utils.Find(servers, func(ms *mcservermanager.MinecraftServer) bool {
		return q.Server == ms.Config.Name
	})
	if index == -1 {
		ws.WriteJSON(errors.ErrorInvalidQuery)
		return
	}
	server := servers[index]
	if err != nil {
		ws.WriteJSON(errors.ErrorInvalidQuery)
		return
	}
	err = ws.WriteJSON(Res{
		Type: TypeLogs,
		Text: server.ReadLogs(),
	})
	if err != nil {
		return
	}
	shouldReturn := false
	go func() {
		for {
			msg := server.GetUpdates()
			err := ws.WriteJSON(Res{
				Type: TypeLogs,
				Text: msg,
			})
			if err != nil {
				return
			}
		}
	}()
	go func() {
		for {
			msg := Req{}
			err := ws.ReadJSON(&msg)
			if err != nil {
				fmt.Println("err ", err)
				shouldReturn = true
				return
			}
			if msg.Type == TypeKeepAlive {
				shouldReturn = false
			}
		}
	}()
	fmt.Println("goroutine opened")
	for {
		time.Sleep(time.Second * 10)
		if shouldReturn {
			fmt.Println("goroutine closed")
			return
		}
		shouldReturn = true
	}

}
