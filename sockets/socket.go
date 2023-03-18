package sockets

import (
	"context"
	"io"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"nhooyr.io/websocket"
)

type Server struct {
}

type Manager struct {
}

func (s *Server) ServeHttp(gc *gin.Context) {
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
		err = handler(r.Context(), c)

		if websocket.CloseStatus(err) == websocket.StatusGoingAway {
			println("closing")
			return
		}
	}

}

func handler(ctx context.Context, c *websocket.Conn) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	println("blocking now...")
	typ, r, err := c.Reader(ctx)
	println("omg message")
	if err != nil {
		println("xd", err.Error())
		return err
	}

	buf := new(strings.Builder)
	_, err = io.Copy(buf, r)
	if err != nil {
		println("ERMG ERR")
	}

	println("clients msg was:", buf.String())
	switch buf.String() {
	case "hi":
		err = c.Write(ctx, typ, []byte("Client said hi"))
		if err != nil {
			println("ERMG ERR")
		}
	case "bye":
		err = c.Write(ctx, typ, []byte("Client said bye"))
		if err != nil {
			println("ERMG ERR")
		}
	default:
		err = c.Write(ctx, typ, []byte("Don't understand what client said"))
		if err != nil {
			println("ERMG ERR")
		}
	}

	//err = c.Write(ctx, typ, []byte("such cool"))

	return err
}
