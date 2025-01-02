package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	cep := "89802250"
	ch1 := make(chan string)
	ch2 := make(chan string)

	fmt.Printf("Buscando cep %s...\n\n", cep)

	go CepBrasilAPI(ch1, cep)
	go CepViaCep(ch2, cep)

	select {
	case data := <-ch1:
		fmt.Printf("Recebido via CepBrasilAPI: %s\n", data)
	case data := <-ch2:
		fmt.Printf("Recebido via CepViaCep: %s\n", data)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout error")
	}
}

func CepBrasilAPI(ch chan string, cep string) {
	resp, error := http.Get("https://brasilapi.com.br/api/cep/v1/" + cep)
	if error != nil {
		return
	}
	defer resp.Body.Close()
	body, error := io.ReadAll(resp.Body)
	if error != nil {
		return
	}
	ch <- string(body)
}

func CepViaCep(ch chan string, cep string) {
	resp, error := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if error != nil {
		return
	}
	defer resp.Body.Close()
	body, error := io.ReadAll(resp.Body)
	if error != nil {
		return
	}

	ch <- string(body)
}
