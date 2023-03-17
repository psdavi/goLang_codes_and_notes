package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const monitoramentos = 5
const intervalo = 1

func main() {

	exibeIntroducao()

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

	response, _ := http.Get(site)

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

/*
UMA CONSTANTE NÂO PODE SER ALTERADA AO LONGO DO CODIGO


Para tal, a cada teste, podemos pedir para o Go esperar um pouco.
Fazemos isso utilizando a função Sleep, do pacote time, passando para ela o
quanto de tempo queremos esperar. Representamos o tempo através de
constantes da própria biblioteca, como Second, Minute, entre outras

Por último, vamos nos livrar de alguns números do nosso código,
e exportá-los para constantes. Por que constantes?
Pois elas não podem ser modificadas.

Os números que queremos atacar é o número 5 que está dentro do for,
que representa o número de monitoramentos, e o número 5 dentro do Sleep,
 que representa o delay do nosso monitoramento.

*/
