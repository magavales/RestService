package main

import (
	utilites "RestService"
	"RestService/pkg/handler"
	"flag"
	"log"
)

func main() {
	flag.Parse()

	router := new(handler.Handler).InitRouter()

	server := new(utilites.Server)
	err := server.InitServer("8080", router)
	if err != nil {
		log.Fatalf("Server can't be opened: %s", err)
	}
}
