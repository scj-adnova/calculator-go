package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/scj-adnova/calculator-go/calculator"
)

var svc = calculator.NewService()

func calculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req calculator.CalculateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	resp := svc.Calculate(req)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func historyHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(svc.GetHistory())
}

func main() {
	http.HandleFunc("/api/calculate", calculateHandler)
	http.HandleFunc("/api/history", historyHandler)

	fmt.Println("Calculator running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
