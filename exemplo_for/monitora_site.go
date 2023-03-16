package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	for {
		resp, err := http.Get("https://portal.ifba.edu.br/")
		if err != nil {
			fmt.Println("Erro ao acessar o site:", err)
		} else {
			fmt.Println("Código de status:", resp.StatusCode)
			resp.Body.Close()
		}
		time.Sleep(time.Minute) // Aguarda um minuto antes de fazer a próxima solicitação
	}
}
