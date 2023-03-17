package main

import (
	"fmt"
	"reflect"
)

func main() {
	chamaSlice()
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
	fmt.Println()

	fmt.Println("*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-")

	fmt.Println()
	fmt.Println("--------------------SLICE-APPEND-----------------------------")
	fmt.Println("Monitorando...")

	nomes = append(nomes, "Davi")
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

Ou seja, o slice dobrou de tamanho quando adicionamos um novo item!
Então, sempre que estourarmos a capacidade máxima do slice,
do array abaixo dele, ele dobra de tamanho.

Logo, o slice nada mais é do que o Go cuidando do array para nós,
pois eles não funcionam de forma diferente.
O slice é um array com algumas coisas abstraídas,
evitando com que nos preocupemos com o tamanho e capacidade do array,
focando somente em trabalhar com os dados.

*/
