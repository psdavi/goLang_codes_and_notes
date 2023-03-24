package contas

import "OOgo/banco/composicaoDeStructs/clientes"

type ContaCorrente struct {
	//letra minuscula o conteudo é acessado apenas pelo arquivo
	//letra maiuscula o conteudo pode ser acessados por outros arquivos
	//Semelhante ao public ou private do java

	// Titular não é mais uma string e sim do tipo Titular
	//pra isso precisa importar do pacote de clientes
	//Se chama composição de Structs

	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64

	// titular       string
	// numeroAgencia int
	// numeroConta   int
	// Saldo         float64
}

// função de 1 retorno
func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque <= c.Saldo && valorDoSaque > 0
	if podeSacar {
		c.Saldo = c.Saldo - valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo Insuficiente"
	}
}

// função de 2 retornos
func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	podeDepositar := valorDoDeposito > 0
	if podeDepositar {
		c.Saldo = c.Saldo + valorDoDeposito
		return "Depósito realizado com sucesso. Saldo:", c.Saldo
	} else {
		return "O valor não é maior que zero", c.Saldo
	}
}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia < c.Saldo && valorDaTransferencia > 0 {
		c.Saldo = c.Saldo - valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}
}
