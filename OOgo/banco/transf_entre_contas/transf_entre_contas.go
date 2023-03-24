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

// função de 2 retornos
func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	podeDepositar := valorDoDeposito > 0
	if podeDepositar {
		c.saldo = c.saldo + valorDoDeposito
		return "Depósito realizado com sucesso. Saldo:", c.saldo
	} else {
		return "O valor não é maior que zero", c.saldo
	}
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia < c.saldo && valorDaTransferencia > 0 {
		c.saldo = c.saldo - valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}
}

func main() {

	contaDoDavi := ContaCorrente{"Davi", 111, 11111, 1000.50}
	contaDoTiti := ContaCorrente{"Titi", 222, 22222, 50000.50}

	status := contaDoTiti.Transferir(40000, &contaDoDavi)

	fmt.Println(status)
	fmt.Println(contaDoDavi)
	fmt.Println(contaDoTiti)

	//contaDoZeca := ContaCorrente{"Zeca", 333, 33333, 7500.60}

}

/*

* aponta pra conta do destinatario
& mostra o endereço da conta do destinatario


*/
