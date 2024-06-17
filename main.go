package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//go:embed static
var static embed.FS

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/units", UnitsHandler)
	router.HandleFunc("/", HomeHandler)
	fmt.Printf("Starting server on port :8080\n")
	log.Fatal(http.ListenAndServe(":8080", router))
}
