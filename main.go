package main

import (
	"log"
	"net/http"

	"morsecode/router"
)

func main() {

	r := router.Router()
	log.Fatal(http.ListenAndServe(":8080", r))
}
