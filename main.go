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

// Define global slice of max capacity 10
var Packages = CreatePackagesSlice()

func main() {
	fmt.Println(Packages)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/units", UnitsHandler)
	router.HandleFunc("/packs", PacksHandler)

	fmt.Printf("Starting server on port :8080\n")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func CreatePackagesSlice() []SPackage {
	slice := make([]SPackage, 5, 10)

	slice[0] = SPackage{1, 250, 0}
	slice[1] = SPackage{2, 500, 0}
	slice[2] = SPackage{3, 1000, 0}
	slice[3] = SPackage{4, 750, 0}
	slice[4] = SPackage{5, 2, 0}

	return slice
}
