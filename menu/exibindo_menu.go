package main

import "fmt"

func main() {
	nome := "Davi"
	idade := 27
	versao := 2.2

	fmt.Println()
	fmt.Println("Olá sr.", nome, "sua idade é", idade)
	fmt.Println("Este programa está na versão", versao)
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
	fmt.Println()

	var comando int
	fmt.Scanf("%d", &comando)
	fmt.Println("O comando escolhido foi", comando)
}

/*
%d representa um número inteiro
%f representa um float
%s representa um string
& significa o endereço da variável que queremos salvar a entrada
&comando recebe o que o usuario digitar no Scanf
*/

/*
& significa o endereço da variável que queremos salvar a entrada,
pois a função Scanf não espera uma variável, e sim o seu endereço, um ponteiro para a variável.
//fmt.Println("O endereço da variável comando é: ", &comando)
*/

/*
var idade int
fmt.Scan(&idade)
*/
