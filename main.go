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

// Define global array of fixed length packages, used later so there can be no more than 10 package sizes
var PackageSizes = [10]SPackage{{1, 100}, {2, 500}, {3, 1000}, {4, 2000}, {5, 5000}, {6, 0}, {7, 0}, {8, 0}, {9, 0}, {10, 0}}

func main() {
	fmt.Println(PackageSizes)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/units", UnitsHandler)
	router.HandleFunc("/packs", PacksHandler)

	fmt.Printf("Starting server on port :8080\n")
	log.Fatal(http.ListenAndServe(":8080", router))
}
