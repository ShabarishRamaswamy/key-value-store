package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/shabarishramaswamy/key-value-store/src/router"
)

var PORT string = ":8529"

func main() {

	router := router.GetNewRouter()

	fmt.Println("Listening to PORT ", PORT)
	log.Fatal(http.ListenAndServe(PORT, router))
}
