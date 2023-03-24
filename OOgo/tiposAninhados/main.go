package main

import (
	"OOgo/tiposAninhados/clientes"
	"OOgo/tiposAninhados/contas"
	"fmt"
)

func main() {

	contaDoDavi := contas.ContaCorrente{
		Titular: clientes.Titular{
			Nome:      "Davi",
			CPF:       "123.123.123-12",
			Profissao: "Pedreiro",
		},
		NumeroAgencia: 123,
		NumeroConta:   123,
		Saldo:         1000}

	fmt.Println(contaDoDavi)
}
