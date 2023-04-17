package main

import (
	"github/justbriang/wsstarter/internal/handlers"
	"log"
	"net/http"
)

func main() {
	//get routes
	routes := routes()

	log.Println("starting channel listener")
	go handlers.ListenToWsChannel()

	log.Println("starting web server at port 8080")
	err := http.ListenAndServe(":8080", routes)
	if err != nil {
		panic(err)
	}

}
