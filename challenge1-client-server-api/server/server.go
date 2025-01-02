package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/andrevfarias/pos-goexpert/challenge1-client-server-api/server/cotacao"
)

func main() {
	http.HandleFunc("/cotacao", HandleCotacao)
	fmt.Println("Server running on port 8080.")
	http.ListenAndServe(":8080", nil)
}

func HandleCotacao(w http.ResponseWriter, r *http.Request) {
	cot, err := cotacao.ObterCotacao()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}

	err = cot.SalvarCotacao()
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cot)
}
