package main

import (

	"databaseserverintegration/routers"
	"fmt"
	"log"
	"net/http"

)

func main() {

	router := routers.Router()
	fmt.Println("Server is starting... http://localhost:4000")
	log.Fatal(http.ListenAndServe(":4000", router))
}
