package sockets

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
)

var (
	upgrader = websocket.Upgrader{}
)

func WsQuote(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	defer ws.Close()

	for {
		// Write
		err := ws.WriteMessage(websocket.TextMessage, []byte("Hello, Client!"))
		if err != nil {
			log.Fatal(err)
		}
		// Read
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", msg)
	}
}
