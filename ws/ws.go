package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
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
	w.Write([]byte("{ \"state\": " + arrayToString(state) + " }"))
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
		log.Fatal(err)
	}
	defer ws.Close()

	clients[ws] = true

	for {
		var stateUpdate StateUpdate
		err := ws.ReadJSON(&stateUpdate)
		if err != nil {
			delete(clients, ws)
			break
		}
		broadcast <- stateUpdate
	}
}

func handleStateUpdates() {
	for {
		stateUpdate := <-broadcast

		state[stateUpdate.Index] = stateUpdate.Value

		for client := range clients {
			err := client.WriteJSON(stateUpdate)
			if err != nil {
				client.Close()
				delete(clients, client)
			}
		}
	}
}
