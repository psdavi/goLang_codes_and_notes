package main

import (
	"fmt"
)

func main() {

	forTradicionalGoLang()
	forGoLang()
}

func forTradicionalGoLang() {
	fmt.Println("Monitorando...")
	sites := []string{"https://portal.ifba.edu.br/", "https://www.youtube.com/", "https://www.casamentos.com.br/"}
	fmt.Println(sites)

	for i := 0; i < len(sites); i++ {
		fmt.Println(sites[i])

	}

	fmt.Println()
	fmt.Println()
	fmt.Println()

}

func forGoLang() {
	fmt.Println("Monitorando...")
	sites := []string{"https://portal.ifba.edu.br/", "https://www.youtube.com/", "https://www.casamentos.com.br/"}
	fmt.Println(sites)

	for i, site := range sites {
		fmt.Println("Estou passando na posicao", i, "do meu slice e essa posição tem o site: ", site)

	}
}

/*

Vimos o for infinito, que faz com o que o código seja repetido para sempre, mas também podemos utilizar o for "tradicional",
onde declaramos uma variável e vamos incrementado-a até o tamanho de itens do slice

Mas há uma forma mais enxuta de fazer isso em Go, utilizando o range.
Ela é como se fosse um operador de iteração do Go, nos dando acesso a cada item
do array, ou do slice, e ele nos retorna dos valores, a posição do item iterado e o
 próprio item daquela posição

*/
