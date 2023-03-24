package main

import (
	"OOgo/banco_completo/clientes"
	"OOgo/banco_completo/contas"
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
