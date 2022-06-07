package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func getTrainLines(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).Encode(sendTrainLinesRequest()); err != nil {
		log.Printf(err.Error())
		panic(err)
	}
}
