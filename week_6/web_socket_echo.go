package week_6

import (
	"GolangBootcamp/common"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var loader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func echoWebSocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := loader.Upgrade(w, r, responseHeder(w, r))
	if err != nil {
		log.Println("error in connection ", err)
		return
	}
	defer conn.Close()
	var u common.Username
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("err in Read JSON", err)
			break
		}
		err = json.Unmarshal(msg, &u)
		if err != nil {
			log.Println("err in unmarshal ", err)
			return
		}

		log.Printf("Received user : %+v\n", u)

		err = conn.WriteJSON(u)
		if err != nil {
			log.Println("err in write JSON", err)
			break
		}
	}

}

func StartWebSocketEchoServer() {
	http.HandleFunc("/ws", echoWebSocketHandler)
	log.Println("WebSocket server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func responseHeder(w http.ResponseWriter, r *http.Request) http.Header {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("web-socket-Created", "Ali")
	return w.Header()
}
