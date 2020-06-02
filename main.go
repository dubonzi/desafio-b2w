package main

import (
	"context"
	"desafio-b2w/conf"
	"desafio-b2w/db"
	"desafio-b2w/logger"
	"desafio-b2w/routes"
	"log"
	"net/http"
	"os"
	"os/signal"
)

const port = ":9080"
const version = "0.1.0"

func main() {
	conf.Load()
	db.DBUri = conf.MongoDBURI()
	db.DBName = conf.MongoDBDatabaseName()
	db.Open()

	server := http.Server{Addr: port, Handler: routes.All()}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	go func() {
		log.Printf("- B2W Planets API v%s -", version)
		log.Printf("- Listening on port %s -", port)
		logger.Fatal("main", "server.ListenAndServe", server.ListenAndServe())
	}()

	<-stop
	log.Println("- Stopping the server -")
	db.Close()
	server.Shutdown(context.Background())
	log.Println("- Server successfully stopped -")
}
