package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Payload struct {
	Event string `json:"event"`
	Data  string `json:"data"`
}

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	var payload Payload
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	fmt.Printf("Evento recibido: %s con datos: %s\n", payload.Event, payload.Data)
	w.WriteHeader(http.StatusOK)
}

func main() {
	http.HandleFunc("/webhook", webhookHandler)
	http.ListenAndServe(":8080", nil)
}
