package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

// função de 1 retorno
func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque <= c.saldo && valorDoSaque > 0
	if podeSacar {
		c.saldo = c.saldo - valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo Insuficiente"
	}
}

//função de 2 retornos

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	podeDepositar := valorDoDeposito > 0
	if podeDepositar {
		c.saldo = c.saldo + valorDoDeposito
		return "Depósito realizado com sucesso. Saldo:", c.saldo
	} else {
		return "O valor não é maior que zero", c.saldo
	}
}

func main() {

	contaDoDavi := ContaCorrente{"Davi", 111, 11111, 1000.50}

	fmt.Println("-------------------------")
	fmt.Println("Saldo atual: ", contaDoDavi.saldo)
	fmt.Println(contaDoDavi.Sacar(200))
	fmt.Println("-------------------------")

	contaDoTiti := ContaCorrente{"Titi", 222, 22222, 50000.50}

	fmt.Println("-------------------------")
	fmt.Println("Saldo atual: ", contaDoTiti.saldo)
	fmt.Println(contaDoTiti.Depositar(200))
	fmt.Println("-------------------------")

	contaDoZeca := ContaCorrente{"Zeca", 333, 33333, 7500.60}

	// OUTRA FORMA DE RETORNAR OS DADOS É USANDO função de multiplos retornos

	fmt.Println("-------------------------")
	fmt.Println("Saldo atual: ", contaDoZeca.saldo)

	msg, valor := contaDoZeca.Depositar(200)
	fmt.Println(msg, valor)
	fmt.Println("-------------------------")

}
