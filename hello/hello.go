package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5

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
			imprimeLogs()
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Nao reconheco esse comando!")
			os.Exit(-1)
		}
	}
}

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
	fmt.Println("")
	sites := leSitesDoArquivo()

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
	}
}

func testaSite(site string) {
	if site == "" {
		fmt.Println("Site inválido: URL vazia.")
		return
	}

	resp, err := http.Get(site)
	tratamentoDeErro(err)
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		fmt.Println("Site", site, "foi carregado com sucesso!")
		registraLog(site, true)
		fmt.Println("")
	} else {
		fmt.Println("Site", site, "esta com problemas! Status Code:", resp.StatusCode)
		registraLog(site, false)
		fmt.Println("")
	}
}

func leSitesDoArquivo() []string {
	var sites []string

	arquivo, err := os.Open("sites.txt")
	tratamentoDeErro(err)
	defer arquivo.Close()

	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		if linha != "" {
			sites = append(sites, linha)
		}
		if err == io.EOF {
			break
		} else if err != nil {
			tratamentoDeErro(err)
			break
		}
	}
	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	tratamentoDeErro(err)
	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + "- online: " + strconv.FormatBool(status) + "\n")
	arquivo.Close()
}

func imprimeLogs() {
	arquivo, err := os.ReadFile("log.txt")
	tratamentoDeErro(err)
	fmt.Println(string(arquivo))
}

func tratamentoDeErro(err error) {
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}
}
