package main

import (
	"flag"
	"fmt"
	"github.com/TutorialEdge/realtime-chat-go-react/db"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

var Router = mux.NewRouter().StrictSlash(true)
var Server = &http.Server{}

func main() {
	flag.Parse()
	db.InitDB()

	ChatHandlers()
	MainHandlers()
	Store_Handlers()
	Server = &http.Server{
		Handler:      Router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println(Server.ListenAndServe())
}
