package main

import "fmt"

func main() {
	nome := "Davi"
	versao := 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")

	var comando int
	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi", comando)

	switch comando {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Saindo do programa...")
	default:
		fmt.Println("Não conheço este comando")
	}
}

/*
switch não usa break
entra na condição
executa
e sai


Para quem vem de outras linguagens de programação, pode estranhar o não uso do break,
ao final do código de cada caso do switch. O break serviria para evitar a execução
do código de mais de um caso, se mais de um caso for atendido.

No Go, ele não possui o break, pois somente um caso pode ser atendido.
O primeiro caso que for atendido, terá o seu código executado e o switch será encerrado.


*/
