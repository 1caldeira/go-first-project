package main

import (
	"fmt"
	"os"
)

func main() {

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

func monitorar() {}

func exibirLog() {}

func exibeMenu() {
	versao := 1.0
	nome := "Gabriel"

	fmt.Println("Olá sr.", nome)
	fmt.Println("Este programa esta na versao:", versao)

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
	site := "https://www.alura.com.br"
	//	resp, _ = http.Get(site)

}
