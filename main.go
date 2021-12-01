package main

import (
	"fmt"
	"log"
	"net/http"
	"openWapi/router"
)

func main() {
	r := router.Router()
	fmt.Println("starting server on port 8081")

	log.Fatal(http.ListenAndServe(":8081", r))
}
