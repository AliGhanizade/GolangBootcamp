package week_6

import (
	"GolangBootcamp/common"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgraderChat = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan common.MessageWebSocket)
var m common.MessageWebSocket


func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgraderChat.Upgrade(w, r, responseHeader(w))
	if err != nil {
		log.Println("Upgrade error:", err)
		return
	}
	defer ws.Close()
	clients[ws] = true

	for {
		_, msg, err := ws.ReadMessage()
		if err != nil {
			log.Println("error in ReadMessage:", err)
			break
		}
		err = json.Unmarshal(msg, &m)
		if err != nil {
			log.Println("err in unmarshal ", err)
			return
		}
		m.DateTime = time.Now().Format(time.DateTime)
		broadcast <- m
		
	}
}

func handleMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				client.Close()
				log.Println("error in Write Message:", err)
			}
			
		}
	}
}

func StartChatServer() {
	http.HandleFunc("/ws", handleConnections)
	go handleMessages()
	log.Println("WebSocket chat server started on :8080")
	http.ListenAndServe(":8080", nil)
}

func responseHeader(w http.ResponseWriter) http.Header{
	w.Header().Set("Content-Type", "application/json")
	return w.Header()
}