package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"teste-b2w/db"
	"teste-b2w/routes"
)

const port = ":9080"
const version = "0.1.0"

func main() {
	db.Open()
	server := http.Server{Addr: port, Handler: routes.All()}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	go func() {
		log.Printf("- B2W Planets API v%s -", version)
		log.Printf("- Listening on port %s -", port)
		log.Fatal("[FATAL] Error while starting http server: ", server.ListenAndServe())
	}()

	<-stop
	log.Println("- Stopping the server -")
	server.Shutdown(context.Background())
	log.Println("- Server successfully stopped -")
}
