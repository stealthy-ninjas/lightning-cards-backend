package sockets

import (
	"context"
	"net/http"
	"time"

	"nhooyr.io/websocket"
)

type Server struct {
	logf func(f string, v ...interface{})
}

func (s Server) ServeHttp(w http.ResponseWriter, r *http.Request) {
	c, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		Subprotocols: []string{"lc.com"},
	})
	if err != nil {
		s.logf("%v", err)
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

func sayHi(ctx context, c *websocket.Conn) error {
	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
}
