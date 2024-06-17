package main

import (
	"encoding/json"
	"net/http"
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
		var unitString UnitRequestJSON
		err := json.NewDecoder(r.Body).Decode(&unitString)
		if err != nil {
			panic(err)
		}

		// Len will not change so fixed len arr is ideal
		var requiredPackages [5]int

		// Convert unit string to int
		units, err := strconv.Atoi(unitString.Value)
		if err != nil {
			panic(err)

		}

		for units > 0 {
			switch {
			case units >= 5000:
				requiredPackages[4]++
				units -= 5000
			case units >= 2000 && units < 5000:
				requiredPackages[3]++
				units -= 2000
			case units >= 1000 && units < 2000:
				requiredPackages[2]++
				units -= 1000
			case units >= 500 && units < 1000:
				requiredPackages[1]++
				units -= 500
			case units < 500:
				requiredPackages[0]++
				units -= 250
			}
		}

		// Create JSON from the required packages array and send as response
		JSONData := &RequiredPackages{Packages: requiredPackages[:]}
		handleJSONwrite(w)
		json.NewEncoder(w).Encode(JSONData)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func PacksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// Return PackageSizes array as JSON
		handleJSONwrite(w)
		json.NewEncoder(w).Encode(PackageSizes)
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
