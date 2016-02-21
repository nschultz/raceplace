package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/echo", echoHandler)
	router.HandleFunc("/broadcast", broadcastHandler)
	router.HandleFunc("/follow", followHandler)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./html/")))
	http.Handle("/", router)
	http.ListenAndServe(":8080", nil)
}

func broadcastHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, welcome to Race Place's broadcast section %s!", r.URL.Path[1:])
}

func followHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, welcome to Race Place's follow section %s!", r.URL.Path[1:])
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		//log.Println(err)
		return
	}

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			return
		}

		printBinary(p)

		err = conn.WriteMessage(messageType, p)
		if err != nil {
			return
		}
	}
}

func printBinary(s []byte) {
	fmt.Printf("Received b:")
	for n := 0; n < len(s); n++ {
		fmt.Printf("%d,", s[n])
	}
	fmt.Printf("\n")
}
