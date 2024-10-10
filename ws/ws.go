package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Allow all connections by default
		return true
	},
}

var stateLength = 10

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan StateUpdate)
var state = make([]int, stateLength)

type StateUpdate struct {
	Index int `json:"index"`
	Value int `json:"value"`
}

func main() {
	http.HandleFunc("/ws", handleConnections)
	http.HandleFunc("/getState", handleGetState)

	go handleStateUpdates()

	log.Println("HTTP server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleGetState(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	stateJSON := "{ \"state\": " + arrayToString(state) + " }"
	_, err := w.Write([]byte(stateJSON))
	if err != nil {
		log.Printf("Error writing state response: %v", err)
	} else {
		log.Printf("State response sent: %s", stateJSON)
	}
}

func arrayToString(arr []int) string {
	result := "["
	for i, v := range arr {
		if i > 0 {
			result += ","
		}
		result += fmt.Sprintf("%d", v)
	}
	result += "]"
	return result
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error upgrading to websocket: %v", err)
		return
	}
	defer func() {
		ws.Close()
		log.Println("WebSocket connection closed")
	}()

	clients[ws] = true
	log.Println("New WebSocket connection established")

	for {
		var stateUpdate StateUpdate
		err := ws.ReadJSON(&stateUpdate)
		if err != nil {
			log.Printf("Error reading JSON from websocket: %v", err)
			delete(clients, ws)
			break
		}
		log.Printf("Received state update: %+v", stateUpdate)
		broadcast <- stateUpdate
	}
}

func handleStateUpdates() {
	for {
		stateUpdate := <-broadcast
		log.Printf("Broadcasting state update: %+v", stateUpdate)

		state[stateUpdate.Index] = stateUpdate.Value

		for client := range clients {
			err := client.WriteJSON(stateUpdate)
			if err != nil {
				log.Printf("Error writing JSON to websocket: %v", err)
				client.Close()
				delete(clients, client)
			} else {
				log.Printf("State update sent to client: %+v", stateUpdate)
			}
		}
	}
}
