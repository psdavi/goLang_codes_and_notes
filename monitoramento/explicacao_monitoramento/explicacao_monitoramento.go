package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"strings"
	"time"
)

// num de monitoramentos
const monitoramentos = 2

// intervalo em segundos para iniciar o proximo
const delay = 5

func main() {

	iniciarMonitoramento()
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	// chama a funcao para ler os sites do arquivo
	sites := leSitesDoArquivo()

	fmt.Println(sites, " sites lidos do arquivoooooooooooooooooooooooooooooooooooooooooooooo")
	fmt.Println(len(sites), "Tamanho do arquivo aaaaaaaaaaaaaaaaaanteeeeeeeeeeeeeeeeeeeeeeeeeeeeesss")

	fmt.Println(reflect.TypeOf(sites), "tipo da vareaaaaaaaaaaaaaaaaaaaaaaaaaveeeelll")

	// o \n gera uma linha em branco a mais no arquivo txt
	// por isso removemos ela com a instrução abaixo
	//se deixarmos essa linha em branco a mais no txt o for vai dar erro de ponteiro
	sites = sites[:len(sites)-1]

	fmt.Println(len(sites), "Tamanho do arquivo depooooooooooooooooooooooooooooooois")

	//for para controlar o numero de monitoramentos desejados
	for i := 0; i < monitoramentos; i++ {
		//for para percorrer o slice de sites até alcancar o tamanho de itens (range)
		for i, site := range sites {

			fmt.Println("Testando site", i, ":", site)
			//para cada site vai ser chamada a função de testar
			testaSite(site)
		}
		//tempo desejado para o intervalo de monitoramentos
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

}

// a funcao recebe um site por parametro e obrigatoriamente tem que informar que o tipo é string
func testaSite(site string) {
	//http.Get é uma função de 2 retornos
	//ela recebe o site pra testar
	//pode retornar resp ou nesse caso optamos por nao trabalhar com o segundo retorno usando _
	resp, _ := http.Get(site)

	//verifica se o status é 200
	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		//se nao for 200 vai mostrar qual o status da requisição
	} else {
		fmt.Println("Site:", site, "esta com problemas. Status Code:", resp.StatusCode)
	}
}

func leSitesDoArquivo() []string {

	//a função informa que o seu retorno é um slice de string

	// aqui é criada a variavel sites que é um slice de strings
	var sites []string

	//aqui a biblioteca os abre um arquivo txt
	// a função os.Open é uma função de 2 retornos
	// um retorno pode ser o erro ou pode ser o arquivo desejado
	arquivo, err := os.Open("sites.txt")

	//se o erro for diferente de nulo
	//vai ser printado o erro com sua caracteristica
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	//biblioteca bufio tem um reader de arquivos, ele vai ler o arquivo que foi carregado anteriormente
	leitor := bufio.NewReader(arquivo)

	for {
		//a funcao ReadString tbm possui 2 retornos
		//a funcao dela é ler uma string até uma quebra de linha \n ou caractere desejado
		linha, err := leitor.ReadString('\n')
		//o \n gera uma linha em branco em cada quebra, então o Trim Space deleta essa linha vazia
		linha = strings.TrimSpace(linha)

		//toda vez que um site é adicionado no txt, ele é adicionado no slice com o append
		sites = append(sites, linha)

		// EOF significa end of file, ultima linha do arquivo txt
		if err == io.EOF {

			fmt.Println(sites)
			//break sai da funcao
			break

		}

	}
	fmt.Println("Saiu do break")

	fmt.Println(sites)

	//retorna o slice com strings de sites
	return sites

}
