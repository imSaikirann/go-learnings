package main 

import (
	"log"

	"encoding/json"
	"net/http"
)


func mains(){
	http.HandleFunc("/ping",pingHandler)
	log.Println("server running ")
	log.Fatal(http.ListenAndServe(":8080",nil))
}

func pingHandler(w http.ResponseWriter , r*http.Request){
	response := map[string]string {
		"message":"pong",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}