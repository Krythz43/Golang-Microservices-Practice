package main

import (
	"fmt"
	"log"
	"net/http"
)

const webPort = "8080"

type Config struct{}

func main() {
	app := Config{}

	log.Printf("Starting broker service on port %s\n", webPort)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	// starting the server

	err := srv.ListenAndServe()

	if err != nil {
		fmt.Println("Error occured: " + err.Error())
		log.Panic(err)
	}
}
