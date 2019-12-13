package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main2() {
	port := ":8080"
	http.HandleFunc("/", HandleDocxRequest)
	fmt.Printf("Server started on port %s", port)
	log.Fatal(http.ListenAndServe(port, nil))
}

func HandleDocxRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	aspDocFields := ASPDocFields{}
	err := decoder.Decode(&aspDocFields)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	encoder := json.NewEncoder(w)
	if err = encoder.Encode(&aspDocFields); err != nil {
		w.Write([]byte(err.Error()))
	}
}
