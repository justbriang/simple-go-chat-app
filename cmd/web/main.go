package main

import (
	"log"
	"net/http"
	"wsstarter/internal/handlers"
)

func main() {
	//get routes
	routes := routes()

	log.Println("starting channel listener")
	go handlers.ListenToWsChannel()

	log.Println("starting web server at port 8080")
	_ = http.ListenAndServe(":8080", routes)

}
