package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	for {
		versao := 1.0
		nome := "Gabriel"
		fmt.Println("Olá sr.", nome)
		fmt.Println("Este programa esta na versao", versao)
		exibeMenu()
		entrada := digitarEntrada()

		switch entrada {
		case 1:
			iniciarMonitoramento()

		case 2:
			fmt.Println("Exibindo logs...")
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Nao reconheco esse comando!")
			os.Exit(-1)
		}
	}
}

func exibirLog() {}

func exibeMenu() {

	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair do programa")
}

func digitarEntrada() int {
	var entrada int
	fmt.Scan(&entrada)
	return entrada
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := []string{"https://httpbin.org/status/200", "https://www.alura.com.br",
		"https://httpbin.org/status/404", "https://www.youtube.com"}

	for i, site := range sites {
		fmt.Println("Testando site", i, ":", site)
		testaSite(site)
	}
}

func testaSite(site string) {
	resp, _ := http.Get(site)
	if resp.StatusCode == 200 {
		fmt.Println("Site", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site", site, "esta com problemas! Status Code:", resp.StatusCode)
	}
}
