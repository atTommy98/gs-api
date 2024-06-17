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

		var requiredPackages [5]int

		// Convert unit string to int
		units, err := strconv.Atoi(unitString.Value)
		if err != nil {
			panic(err)

		}

		for units > 0 {
			// Units greater than 5000, add a pack of 5000
			if units >= 5000 {
				requiredPackages[4]++
				units = units - 5000
				// Units greater than 2000, but less than 5000, add a pack of 2000
			} else if units >= 2000 && units < 5000 {
				requiredPackages[3]++
				units = units - 2000
				// Units greater than 1000, but less than 2000, add a pack of 1000
			} else if units >= 1000 && units < 2000 {
				requiredPackages[2]++
				units = units - 1000
				// Units greater than 500, but less than 1000, add a pack of 500
			} else if units >= 500 && units < 1000 {
				requiredPackages[1]++
				units = units - 500
				// Units less than 500, add a pack of 250
			} else {
				requiredPackages[0]++
				units = units - 250
			}
		}

		// Create JSON from the required packages array and send as response
		JSONData := &RequiredPackages{Packages: requiredPackages[:]}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(JSONData)
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
