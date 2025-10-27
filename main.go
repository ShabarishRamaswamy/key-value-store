package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/shabarishramaswamy/key-value-store/src/commonMethods"
	"github.com/shabarishramaswamy/key-value-store/src/databases"
	"github.com/shabarishramaswamy/key-value-store/src/models"
	"github.com/shabarishramaswamy/key-value-store/src/router"
)

var PORT string = ":8529"

func main() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	db, err := databases.RunInitDBLoop()
	if err != nil {
		return
	}

	ctx := context.WithValue(context.Background(), models.GlobalKVStoreName, commonMethods.GetNewKVPair())
	router := router.GetNewRouter(ctx, db)

	fmt.Println("Listening to PORT ", PORT)
	go func() {
		log.Fatal(http.ListenAndServe(PORT, router))
	}()
	sig := <-quit

	log.Printf("Main: Received signal: %v. Initiating graceful shutdown...", sig)
	db.Close()
}
