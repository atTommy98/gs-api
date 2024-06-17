package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	p, err := static.ReadFile("static/index.html")
	if err != nil {
		panic(err)
	}
	w.Write(p)
}

func UnitsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var unitString UnitPostJSON
		err := json.NewDecoder(r.Body).Decode(&unitString)
		if err != nil {
			panic(err)
		}

		// Create 'units' int variable
		units, err := strconv.Atoi(unitString.Value)
		if err != nil {
			panic(err)

		}

		// Create copy of packages and order by size
		PackagesAscending := make([]SPackage, len(Packages))
		copy(PackagesAscending, Packages)
		sort.Slice(PackagesAscending, func(i, j int) bool {
			return PackagesAscending[i].Size < PackagesAscending[j].Size
		})

		// Reset required packages
		for i := 0; i < len(Packages); i++ {
			Packages[i].Required = 0
		}

		// Update based on size vs units left using id
		for units > 0 {
			for i := 0; i < len(PackagesAscending); i++ {
				// If current iteration is the largest pack and units are greater than it, just add largest size pack
				if i+1 == len(PackagesAscending) && units >= PackagesAscending[i].Size {
					fmt.Printf("The largest pack size available is %v, incrementing required pack\n", PackagesAscending[i].Size)
					for j := 0; j < len(Packages); j++ {
						if PackagesAscending[i].Id == Packages[j].Id {
							Packages[j].Required += 1
							units -= PackagesAscending[i].Size
						}
					}
					// If units are greater than current size, move to the next size
				} else if units > PackagesAscending[i].Size {
					fmt.Printf("Units are greater than size %v, moving on\n", PackagesAscending[i].Size)
					continue
					// If units are less than current size, add the size below
				} else if units <= PackagesAscending[i].Size {
					fmt.Printf("Units are smaller than size %v, incrementing required pack\n", PackagesAscending[i].Size)
					for j := 0; j < len(Packages); j++ {
						if i == 0 || units == PackagesAscending[i].Size {
							if PackagesAscending[i].Id == Packages[j].Id {
								Packages[j].Required += 1
								units -= PackagesAscending[i].Size
							}
						} else {
							if PackagesAscending[i-1].Id == Packages[j].Id {
								Packages[j].Required += 1
								units -= PackagesAscending[i-1].Size
							}
						}
					}
					break
				}
				fmt.Printf("Units left are %v\n", units)
			}
		}

		// Create JSON from the required packages array and send as response
		handleJSONwrite(w)
		json.NewEncoder(w).Encode(Packages)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func PacksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// Return Packages array as JSON
		handleJSONwrite(w)
		json.NewEncoder(w).Encode(Packages)
	case "POST":
		// Do something
	case "PUT":
		// Do something
	case "DELETE":
		// Do something
	default:
		w.WriteHeader(http.StatusBadRequest)
	}

}

func handleJSONwrite(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusCreated)
}
