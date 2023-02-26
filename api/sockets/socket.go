package sockets

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"nhooyr.io/websocket"
)

type Server struct {
	Logf func(f string, v ...interface{})
}

func (s Server) ServeHttp(gc *gin.Context) {
	w := gc.Writer
	r := gc.Request
	c, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		Subprotocols:   []string{"lc.com"},
		OriginPatterns: []string{"*"},
	})
	if err != nil {
		s.Logf("%v", err)
		return
	}
	defer c.Close(websocket.StatusInternalError, "The sky is pink")

	for {
		err = sayHi(r.Context(), c)
		if websocket.CloseStatus(err) == websocket.StatusNormalClosure {
			return
		}
	}
}

func sayHi(ctx context.Context, c *websocket.Conn) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	typ, r, err := c.Reader(ctx)
	if err != nil {
		return err
	}

	w, err := c.Writer(ctx, typ)
	if err != nil {
		return err
	}

	_, err = io.Copy(w, r)
	if err != nil {
		return fmt.Errorf("failed to io copy: %w", err)
	}

	err = w.Close()
	return err
}
