package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins
	},
}

func main() {
	http.HandleFunc("/", handleConnection)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func handleConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer func(conn *websocket.Conn) {
		err := conn.Close()
		if err != nil {
			log.Printf("Error while closing connection: %v", err)
		}
	}(conn)

	geoFenceClient := NewGeoFenceClient(grpcClientConn) // Create your gRPC client here

	webSocketClient := &WebSocketGeoFenceClient{wsConn: conn}
	streamCopier := &StreamCopier{src: geoFenceClient, dest: webSocketClient}

	go streamCopier.StartCopying()
	select {}
}
