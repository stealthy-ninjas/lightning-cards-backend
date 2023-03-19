package sockets

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
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
	fmt.Println("THING IS", s.rooms)
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
			println("closing")
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
		s.rooms[rId].Players[p.Username] = p
		s.rooms[rId].Running = true
		println(rId)

		// todo(): alert other players of joining
		for _, p := range s.rooms[rId].Players {
			p.Ws_connection.Write(ctx, typ, []byte("hi hello"))
		}

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

	// buf := new(strings.Builder)
	// _, err = io.Copy(buf, r)
	// if err != nil {
	// 	println("ERMG ERR")
	// }

	// println("clients msg was:", buf.String())
	// switch buf.String() {
	// case "hi":
	// 	err = c.Write(ctx, typ, []byte("Client said hi"))
	// 	if err != nil {
	// 		println("ERMG ERR")
	// 	}
	// case "bye":
	// 	err = c.Write(ctx, typ, []byte("Client said bye"))
	// 	if err != nil {
	// 		println("ERMG ERR")
	// 	}
	// default:
	// 	err = c.Write(ctx, typ, []byte("Don't understand what client said"))
	// 	if err != nil {
	// 		println("ERMG ERR")
	// 	}
	// }

	// //err = c.Write(ctx, typ, []byte("such cool"))

	// return err
}

func jsonExtract(r io.Reader) *models.SocketMessage {
	// buf := new(strings.Builder)
	// _, err := io.Copy(buf, r)
	// if err != nil {
	// 	println("ERMG ERR")
	// }

	// println("clients msg was:", buf.String())

	// j := &models.SocketMessage{}

	bytesBuf := &bytes.Buffer{}
	n, err := io.Copy(bytesBuf, r)
	println(n)
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
