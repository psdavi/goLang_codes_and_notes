package main

// pacote de formatação, formatar e imprimir
import (
	"fmt"
	"reflect"
)

// Para saber se o Go conseguir inferir corretamente o tipo das variáveis, podemos descobri-los, importando o pacote reflect

func main() {

	var nome string = "Davi"
	var versao float32 = 2.0
	var idade int = 27

	var salario = 100.50
	var carros = 3

	nome_2 := "Davi"
	versao_2 := 2.0
	idade_2 := 27

	fmt.Println("- - - - - - - - - - - - - - - - - - - - -")
	fmt.Println("Olá senhor", nome)
	fmt.Println("Olá senhor", nome, "Sua idade é", idade)
	fmt.Println("Este programa está na versão", versao)
	fmt.Println("- - - - - - - - - - - - - - - - - - - - -")
	fmt.Println("O tipo da variável salario é", reflect.TypeOf(salario))
	fmt.Println("O tipo da variável salario é", reflect.TypeOf(carros))
	fmt.Println("- - - - - - - - - - - - - - - - - - - - -")
	fmt.Println("Olá senhor", nome_2)
	fmt.Println("Olá senhor", nome_2, "Sua idade é", idade_2)
	fmt.Println("Este programa está na versão", versao_2)
	fmt.Println("- - - - - - - - - - - - - - - - - - - - -")

}

/*

	* quando declara uma variavel vazia ela recebe automaticamente:

	string = vazio
	float = 0.0
	int = 0

	* quando uma variavel é declarada mas não é usada, o programa go não compila
	* concatenação com ,

*/

/*

Alternativa correta! O tipo de variável float no Go não existe,
porque ele possui tamanhos, de 32 ou 64 bits, que devem ser especificados
 no momento da sua declaração. Se a variável for de 32 bits, o seu tipo será
 float32, mas se for de 64 bits, o seu tipo será float64.

*/

/*

Para deixar o nosso código mais limpo ainda, podemos remover a palavra var das variáveis.
 Podemos fazer isso pois o Go possui um segundo operador de atribuição de variáveis, um mais
  "curto", que é o :=. Quando utilizamos esse operador, estamos dizendo ao Go que estamos declarando
  uma variável e atribuindo um valor a ela:

*/
