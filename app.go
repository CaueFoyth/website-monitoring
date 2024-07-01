package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {
	
	exibeIntroducao()
	for{
		menu()
		comando := leComando()

		switch comando {
			case 1:
				iniciarMonitoramento()
			case 2:
				imprimeLogs()
			case 0:
				fmt.Println("Saindo do programa...")
				os.Exit(0)
			default:
				fmt.Println("Opcão inválida")
				os.Exit(-1)
		}	
	}
}

func exibeIntroducao() {
	fmt.Println("Olá, qual é o seu nome?")
	var nome string 
	fmt.Scan(&nome)
	
	versao := 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
}

func menu() {
	fmt.Println("1- Iniciar monitoramento")
	fmt.Println("2- Exibir logs")
	fmt.Println("0- Sair do programa")
}

func leComando() int {
	var comandoLido int
	//& = endereço de memória / ponteiro
	fmt.Scan(&comandoLido)

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Iniciando monitoramento...")
	sites := leSitesDoArquivo()
	
	for i := 0; i < monitoramentos; i++ {
		for _, site := range sites{
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}	

	fmt.Println("")
}

func testaSite(site string){
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:",err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "está sem resposta. Status:", resp.StatusCode)
		registraLog(site, false)
	}
}

func leSitesDoArquivo() []string {
	var sites []string
	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)
	
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
		if err == io.EOF {
			break
		}
	}

	arquivo.Close()
	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")
	arquivo.Close()
}

func imprimeLogs() {

	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	fmt.Println(string(arquivo) + "\n")

}