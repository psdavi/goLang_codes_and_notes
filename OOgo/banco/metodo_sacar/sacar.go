package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque <= c.saldo && valorDoSaque > 0
	if podeSacar {
		c.saldo = c.saldo - valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo Insuficiente"
	}

}

func main() {

	contaDoZeca := ContaCorrente{"Zeca", 333, 55555, 100.50}

	contaDoDavi := ContaCorrente{"Davi", 333, 55555, 200.50}

	// ------------------- 2 tipos de tratamento de int e float ---------------------------
	// . no final de um numero passa ele pra float
	saque := 50.

	fmt.Println(contaDoZeca.saldo)

	//da errado pq saque é int e saldo é float
	//contaDoZeca.saldo = contaDoZeca.saldo - saque
	contaDoZeca.saldo = contaDoZeca.saldo - saque

	fmt.Println(contaDoZeca.saldo)

	fmt.Println("-----------------------------------------------------")

	//saque 2 esta do tipo int e é convertido pra float abaixo

	fmt.Println(contaDoDavi.saldo)

	saque2 := 50

	contaDoDavi.saldo = contaDoDavi.saldo - float64(saque2)

	fmt.Println(contaDoDavi.saldo)

	fmt.Println(contaDoDavi.Sacar(200))
	//fmt.Println(contaDoDavi.Sacar(151))

}

/*

Para referenciarmos esse ponteiro no momento da criação do tipo, podemos colocar entre parênteses logo após func e antes de Sacar()
a inscrição (c *ContaCorrente), o que significa que quando a função for chamada, o código apontará para a conta que chama. Nesse caso,
quando chamarmos a função, não precisaremos especificar que nos tratamos da conta de um cliente ou outro.

Nesse caso, se a conta que estiver chamando a função tiver saldo, será possível sacar. Criaremos uma condicional if
 para fazer a verificação. Poderemos sacar se for verdadeiro que valorDoSaque é menor do que saldo. Se podemos sacar colocaremos c.saldo
 no corpo do if. Podíamos escrever conta.saldo, se escrevêssemos (conta *ContaCorrente) , mas por uma questão da linguagem Go, sempre utilizamos
 a primeira letra do nome no nosso ponteiro dentro da função.

*/
