package main

import (
	"fmt"
	"reflect"
)

func main() {
	chamaArray()
	chamaSlice()
}

func chamaArray() {

	fmt.Println()
	fmt.Println("--------------------ARRAY------------------------------")
	fmt.Println("Monitorando...")
	var sites [3]string
	sites[0] = "www.google.com"
	sites[1] = "www.alura.com.br"
	sites[2] = "www.amazon.com"

	fmt.Println(sites)
	fmt.Println("O tipo da variavel é: ", reflect.TypeOf(sites))
	fmt.Println("O meu array tem", len(sites), "itens")
	fmt.Println("O meu array tem capacidade para", cap(sites), "itens")
	fmt.Println("--------------------------------------------------")
	fmt.Println()

}

func chamaSlice() {

	fmt.Println()
	fmt.Println("--------------------SLICE------------------------------")
	fmt.Println("Monitorando...")

	nomes := []string{"Douglas", "Daniel", "Bernardo"}
	fmt.Println(nomes)
	fmt.Println("O meu slice tem", len(nomes), "itens")
	fmt.Println("O tipo da variavel é: ", reflect.TypeOf(nomes))
	fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens")
	fmt.Println("--------------------------------------------------")
}

/*
QUANDO UM ARRAY É DECLARADO COM UM TAMANHO MAIOR DO QUE SEUS ITENS,
OS ITENS QUE NÂO EXISTEM SÂO IMPRESSOS EM BRANCO MAS CONTINUA OCUPANDO OS ESPAÇOS DO ARRAY
*/

/*
UM SLICE È FEITO EM CIMA DE UM ARRAY
*/
