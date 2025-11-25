package main

import (
	"chatroom/models"
	"chatroom/service"
	"chatroom/transport"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	broadcastChan := make(chan models.Message)

	stateService := service.NewStateService(broadcastChan)

	go stateService.HandleMessageLoop()

	wsHandler := transport.NewWebsocketHandler(stateService)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/ws", wsHandler.HandleConnections)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server started on :%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
