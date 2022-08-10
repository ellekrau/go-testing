package main

import (
	"go-testing/database"
	"go-testing/server"
	"log"
)

func main() {
	if err := database.StartDB(); err != nil {
		log.Fatalln(err)
	}
	server.StartServer()
}
