package main

import (
	"fmt"
	"log"

	"github.com/andrefarias66/pos-goexpert/desafios/desafio1-client-server-api/client/cotacao"
)

func main() {
	cot, err := cotacao.ObterCotacao()
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("Cotação do dólar: R$", cot.Bid)

	err = cot.SalvarTxt("cotacao.txt")
	if err != nil {
		log.Println(err)
	}
}
