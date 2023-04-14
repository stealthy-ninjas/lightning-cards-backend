package sockets

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"stealthy-ninjas/lightning-cards/models"

	"github.com/gin-gonic/gin"
	"nhooyr.io/websocket"
)

type service struct {
	rooms models.Rooms
}

func NewService(rooms models.Rooms) *service {
	return &service{rooms: rooms}
}

func (s *service) ServeHttp(gc *gin.Context) {
	w := gc.Writer
	r := gc.Request
	c, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		Subprotocols:   []string{"lc"},
		OriginPatterns: []string{"*"},
	})

	if err != nil {
		log.Println(err)
		return
	}

	defer c.Close(websocket.StatusInternalError, "Connection closing...")

	for {
		err = s.handler(r.Context(), c)

		if websocket.CloseStatus(err) == websocket.StatusGoingAway {
			log.Println("closing")
			return
		}
	}

}

func (s *service) handler(ctx context.Context, c *websocket.Conn) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	typ, r, err := c.Reader(ctx)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	socketMsg := jsonExtract(r)

	switch socketMsg.Type {
	case "join":
		p := models.Player{
			Username:      socketMsg.Body.(map[string]interface{})["userName"].(string),
			Ws_connection: c,
		}
		rId := socketMsg.Body.(map[string]interface{})["roomId"].(string)

		// alert other players that joining player has joined
		joinEvent := &models.SocketMessage{Type: "new join", Body: p}
		marshalledEvent, _ := json.Marshal(joinEvent)
		for _, p := range s.rooms[rId].Players {
			p.Ws_connection.Write(ctx, typ, marshalledEvent)
		}

		// add this player to the room
		s.rooms[rId].Players[p.Username] = p

		// send list of players currently in room to joining player
		lop, err := json.Marshal(s.rooms[rId].Players)
		if err != nil {
			log.Println(err)
		}

		err = c.Write(ctx, typ, lop)
		if err != nil {
			log.Println(err)
		}
	}

	return err
}

func jsonExtract(r io.Reader) *models.SocketMessage {
	bytesBuf := &bytes.Buffer{}
	_, err := io.Copy(bytesBuf, r)
	if err != nil {
		log.Println(err)
	}

	sm := &models.SocketMessage{}
	err = json.Unmarshal(bytesBuf.Bytes(), sm)
	if err != nil {
		log.Println(err)
	}
	return sm
}
