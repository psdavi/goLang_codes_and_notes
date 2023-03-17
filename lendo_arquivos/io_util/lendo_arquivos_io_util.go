package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const monitoramentos = 5
const intervalo = 1

func main() {

	exibeIntroducao()
	leSitesDoArquivo()
	leSitesDoArquivoIoUtil()

	for {
		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciaMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}

	}
}

func exibeIntroducao() {
	fmt.Println()
	nome := "Davi"
	versao := 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)

	fmt.Println()
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)
	return comandoLido
}

func exibeMenu() {
	fmt.Println()
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
	fmt.Println()
}

func iniciaMonitoramento() {
	fmt.Println("Monitorando...")

	sites := []string{"https://portal.ifba.edu.br/", "https://www.youtube.com/", "https://www.casamentos.com.br/"}
	fmt.Println(sites)

	for i := 0; i < monitoramentos; i++ {
		fmt.Println()
		fmt.Println("----MONITORANDO VEZ NÚMERO: ", i, "----")
		fmt.Println()
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(intervalo * time.Second)
	}
}

func testaSite(site string) {

	response, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if response.StatusCode == 200 {
		fmt.Println()
		fmt.Println("Status code:", response.StatusCode, ": O site", site, "foi carregado com sucesso!")
		fmt.Println()
	} else {
		fmt.Println()
		fmt.Println("Houve algum problema ao carregar o site! Status Code:", response.StatusCode)
		fmt.Println()
	}
}

func leSitesDoArquivo() []string {

	var sites []string
	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(arquivo)

	fmt.Println()

	return sites

}

func leSitesDoArquivoIoUtil() []string {

	var sites []string
	arquivo, err := ioutil.ReadFile("sites.txt")
	//imprime array de bytes

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(string(arquivo))
	//converte o array de bytes pra string

	return sites

}

/*
Através do pacote io/ioutil, chamamos a função ReadFile, para ler o arquivo
Essa função nos retorna o arquivo em um array de bytes, então basta
convertê-los para uma string através da função string

*/
