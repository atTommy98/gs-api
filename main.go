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

var PackageSizes = [10]int{250, 500, 1000, 2000, 5000, 0, 0, 0, 0, 0}

func main() {
	fmt.Println(PackageSizes)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/units", UnitsHandler)
	router.HandleFunc("/packs", PacksHandler)

	fmt.Printf("Starting server on port :8080\n")
	log.Fatal(http.ListenAndServe(":8080", router))
}
