package main

import (
	"context"
	"net/http"

	_ "github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/log"
	"github.com/coder/websocket"
)

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Accept(w, r, &websocket.AcceptOptions{
		InsecureSkipVerify: true, // Allow all origins for development
	})
	if err != nil {
		log.Error("Accept error:", err)
		return
	}
	defer conn.Close(websocket.StatusNormalClosure, "")

	ctx := context.Background()

	for {
		messageType, message, err := conn.Read(ctx)
		if err != nil {
			log.Error("Read error:", err)
			break
		}
		log.Printf("Received: %s", message)

		// Echo the message back
		err = conn.Write(ctx, messageType, message)
		if err != nil {
			log.Error("Write error:", err)
			break
		}
	}
}

func main() {
	http.HandleFunc("/ws", handleWebSocket)
	log.Info("WebSocket server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
