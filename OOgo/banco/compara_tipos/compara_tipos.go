package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {

	var contaDoDavi ContaCorrente = ContaCorrente{}
	contaDoTiti := ContaCorrente{}
	contaDoMicael := ContaCorrente{titular: "Micael", numeroAgencia: 123, numeroConta: 12345, saldo: 500.50}
	contaDoZeca := ContaCorrente{"Zeca", 333, 55555, 100.50}

	fmt.Println(contaDoDavi)
	fmt.Println(contaDoTiti)
	fmt.Println(contaDoMicael)
	fmt.Println(contaDoZeca)

	fmt.Println()
	fmt.Println("------------------------COMO JAVA E C++ -------------------------------")
	fmt.Println()

	var contaDoVinicius *ContaCorrente
	contaDoVinicius = new(ContaCorrente)
	contaDoVinicius.titular = "Vinicius"
	contaDoVinicius.saldo = 50000

	var contaDoVinicius2 *ContaCorrente
	contaDoVinicius2 = new(ContaCorrente)
	contaDoVinicius.titular = "Vinicius"
	contaDoVinicius.saldo = 50000

	//Igual a false pq embora o conteudo seja o msm, estao apontando pra endereços diferentes na memoria
	fmt.Println(contaDoVinicius == contaDoVinicius2)
	fmt.Println(*contaDoVinicius)
	fmt.Println(contaDoVinicius)

	//----------------------------------------------------------------------------------------------------------

	contaDoJoao := ContaCorrente{"Joao", 333, 55555, 100.50}
	contaDoJoao2 := ContaCorrente{"Joao", 333, 55555, 100.50}

	//Igual a true pq esta comparando o conteudo armazenado no mesmo espaco de memoria
	fmt.Println(contaDoJoao == contaDoJoao2)
	fmt.Println(contaDoJoao)
	fmt.Println(contaDoJoao2)

	//-------------------------------------------------------------------------------------------------------------

	var contaDaMaria *ContaCorrente
	contaDaMaria = new(ContaCorrente)
	contaDaMaria.titular = "Vinicius"
	contaDaMaria.saldo = 50000

	var contaDaMaria2 *ContaCorrente
	contaDaMaria2 = new(ContaCorrente)
	contaDaMaria2.titular = "Vinicius"
	contaDaMaria2.saldo = 50000

	fmt.Println("------- MARIA ----------")
	fmt.Println("-------conteudo igual mas endereços diferentes---- conteudo não é igual------")
	fmt.Println(contaDaMaria == contaDaMaria2)
	fmt.Println(&contaDaMaria)
	fmt.Println(&contaDaMaria2)

	fmt.Println("------- MARIA ----------")
	fmt.Println("-------comparando apenas o conteudo sem considerar o endereço---conteudo é igual-------")
	fmt.Println(*contaDaMaria == *contaDaMaria2)
	fmt.Println(&contaDaMaria)
	fmt.Println(&contaDaMaria2)

}
