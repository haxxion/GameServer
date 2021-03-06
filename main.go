package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// WebSocket
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "/mnt/e/Programming/Go/GameServer/html/home.html")
}

func main() {
	fmt.Println("Haxxion.xyz Online :)")

	// Home
	http.HandleFunc("/", rootHandler)

	// Multiplayer
	// Connect4
	http.HandleFunc("/connect4/", connect4Handler)
	http.HandleFunc("/connect4echo", connect4EchoHandler)

	// Singleplayer
	// FloatySquare
	http.HandleFunc("/floatysquare/", floatySquareHandler)
	http.HandleFunc("/floatysquareecho", floatySquareEchoHandler)
	// Snake
	http.HandleFunc("/snake/", snakeHandler)
	http.HandleFunc("/snakeecho", snakeEchoHandler)

	// Cellular Automata
	http.HandleFunc("/gameoflife/", gameOfLifeHandler)
	http.HandleFunc("/langtonsant/", langtonsAntHandler)
	http.HandleFunc("/briansbrain/", briansBrainHandler)
	http.HandleFunc("/wireworld/", wireWorldHandler)
	http.HandleFunc("/sandpiles/", sandPilesHandler)

	log.Fatal(http.ListenAndServe(":80", nil))
}
