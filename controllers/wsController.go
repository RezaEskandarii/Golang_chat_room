package controllers

import (
	"../models"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

type WsController struct {
	Echo *echo.Echo
}

var (
	upgrader = websocket.Upgrader{
		// default buffer size fo reading messages
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

var (
	connectedClients map[*websocket.Conn]bool
	broadcastMessage chan models.Message
)

func (w *WsController) Init() {
	connectedClients = make(map[*websocket.Conn]bool)
	broadcastMessage = make(chan models.Message)

	w.Echo.Any("/ws", handleConnections)
	go broadcastMessages()
}

func handleConnections(c echo.Context) error {
	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		log.Fatal(err)
	}
	defer ws.Close()

	connectedClients[ws] = true

	for {
		var message models.Message

		err := ws.ReadJSON(&message)
		if err != nil {
			fmt.Println(err.Error())
			delete(connectedClients, ws)
			break
		}
		broadcastMessage <- message

	}
	return nil
}

func broadcastMessages() {
	for {
		message := <-broadcastMessage
		for client := range connectedClients {
			err := client.WriteJSON(message)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(connectedClients, client)
			}
		}
	}
}
