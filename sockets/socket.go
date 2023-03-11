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

func (s Server) ServeHttp(gc *gin.Context) {
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
		err = sayHi(r.Context(), c)

		if websocket.CloseStatus(err) == websocket.StatusGoingAway {
			println("closing")
			return
		}
	}

}

func sayHi(ctx context.Context, c *websocket.Conn) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	println("blocking now...")
	typ, r, err := c.Reader(ctx)
	println("omg message")
	if err != nil {
		println("xd", err.Error())
		return err
	}

	err = c.Write(ctx, typ, []byte("such cool"))
	if err != nil {
		println("ERMG ERR")
	}
	buf := new(strings.Builder)
	_, err = io.Copy(buf, r)
	println("clients msg was: ", buf.String())
	if err != nil {
		println("ERMG ERR")
	}

	_, err = r.Read([]byte{})
	return err
}
