package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/shabarishramaswamy/key-value-store/src/commonMethods"
	"github.com/shabarishramaswamy/key-value-store/src/databases"
	"github.com/shabarishramaswamy/key-value-store/src/models"
	"github.com/shabarishramaswamy/key-value-store/src/router"
)

var PORT string = ":8529"

func main() {
	db, err := databases.RunInitDBLoop()
	if err != nil {
		return
	}

	ctx := context.WithValue(context.Background(), models.GlobalKVStoreName, commonMethods.GetNewKVPair())
	router := router.GetNewRouter(ctx, db)

	fmt.Println("Listening to PORT ", PORT)
	log.Fatal(http.ListenAndServe(PORT, router))
}
