package main

import "fmt"

func main() {

	fmt.Println("Monitorando...")
	var sites [4]string
	sites[0] = "www.google.com"
	sites[1] = "www.alura.com.br"
	sites[2] = "www.amazon.com"

	fmt.Println(sites[3])

}

/*
o fato do array ter um tamanho fixo,
nos limita um pouco, pois se quisermos adicionar 5 itens no array,
teremos que alterar o seu tamanho na sua declaração.
Por isso, em Go, geralmente não trabalhamos com arrays,
e sim com uma outra estrutura de dados, chamada Slice,
que funciona em cima do array, mas não tem tamanho fixo.
*/
