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

	fmt.Println(contaDoVinicius)
	fmt.Println(*contaDoVinicius)

}
