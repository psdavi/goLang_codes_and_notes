package main

import (
	"fmt"
)

func main() {

	nome, idade, _ := devolveNomeEIdade()
	fmt.Println("Meu nome é ", nome, "e tenho", idade, "anos.")

	fmt.Println()

	_, empresa := devolveCargoEEmpresa()
	fmt.Println("Trabalho na empresa", empresa)

}

func devolveNomeEIdade() (string, int, bool) {
	nome := "Davi"
	idade := 27
	tamanho := true
	return nome, idade, tamanho
}

func devolveCargoEEmpresa() (string, string) {
	cargo := "Estagiário Desenvolvedor"
	empresa := "Veri"
	return cargo, empresa
}

/*
Quando não queremos saber de um dos retornos, queremos ignorá-lo,
nós utilizamos o operador de identificador em branco (_):

*/
