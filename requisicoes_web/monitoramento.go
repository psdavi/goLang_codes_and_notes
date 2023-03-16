package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	exibeIntroducao()
	exibeMenu()
	comando := leComando()

	switch comando {
	case 1:
		iniciaMonitoramento()
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Saindo do programa...")
		os.Exit(0)
	default:
		fmt.Println("Não conheço este comando")
		os.Exit(-1)
	}
}

func exibeIntroducao() {
	nome := "Davi"
	versao := 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)
	return comandoLido
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func iniciaMonitoramento() {
	fmt.Println("Monitorando...")
	site := "https://portal.ifba.edu.br/"
	http.Get(site)

}

/*
pacote para fazer requisições web "net/http"
"net/http" é pacote mais específico à nossa necessidade.
Já que nele temos funções para realizar requisições Get e Post.
Que pela própria definição dele tem como objetivo fornecer a
implementações de cliente e servidor HTTP.

É importante saber que temos vários subdiretórios dentro do "net".
Se quiséssemos fazer um envio de email poderíamos usar o "net/smtp".

existem funções no Go que retornam mais de um valor e a Get é uma delas,
além da resposta, ela também retorna um possível erro que possa ter acontecido
na requisição:

Como referência para sabermos os pacotes da linguagem temos: https://golang.org/pkg/

*/
