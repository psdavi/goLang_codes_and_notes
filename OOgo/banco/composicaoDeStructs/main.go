package main

//importa o caminho onde está o arquivo
import (
	"OOgo/banco/pacotesEVisibilidade/contas"
	"fmt"
)

func main() {

	//para acessar, basta usar o nome do pacote e usar .
	contaDoDavi := contas.ContaCorrente{"Davi", 111, 11111, 1000.50}
	contaDoTiti := contas.ContaCorrente{"Titi", 222, 22222, 50000.50}

	status := contaDoTiti.Transferir(40000, &contaDoDavi)

	fmt.Println(status)
	fmt.Println(contaDoDavi)
	fmt.Println(contaDoTiti)

}
