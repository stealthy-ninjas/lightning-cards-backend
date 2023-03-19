package models

import "nhooyr.io/websocket"

type Rooms map[string]*Room
type Room struct {
	Players map[string]Player
	Running bool
}
type Player struct {
	Ws_connection *websocket.Conn
	Username      string
}
