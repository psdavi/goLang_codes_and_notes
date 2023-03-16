package main

import "fmt"

func main() {
	nome := "Davi"
	versao := 3.3
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")

	var comando int
	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi", comando)

	if comando == 1 {
		fmt.Println("Iniciando monitoramento...")
	} else if comando == 2 {
		fmt.Println("Exibindo logs...")
	} else if comando == 0 {
		fmt.Println("Saindo do programa...")
	} else {
		fmt.Println("Não conheço esse comando")
	}
}

/*
if comando{

}

não funciona no GO
precisa ter uma verificação de True ou False
maiorDeIdade := true
*/
