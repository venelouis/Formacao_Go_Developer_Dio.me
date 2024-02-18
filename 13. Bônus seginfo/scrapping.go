//Scraping Web
//usa biblioteca colly
//extrair dados de um site
//como ele faz: acessa o www pelo protocolo http
//serve: análise de dados, gerar leads
//para Seg Inf: Levantamento de informações para um possível alvo

package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

func main() {
	fName := "data.csv"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Não foi possível criar o arquivo, err: %q", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	c := colly.NewCollector()
	c.OnHTML("table#customers", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			writer.Write([]string{
				el.ChildText("td:nth-child(1)"),
				el.ChildText("td:nth-child(2)"),
				el.ChildText("td:nth-child(3)"),
			})
		})
		fmt.Println("Varredura Completa")
	})
	c.Visit("https://www.w3schools.com/html/html_tables.asp")
}

/*
scrapping.go:16:2: no required module provides package github.com/gocolly/colly: go.mod file not found in current directory or any parent directory; see 'go help modules'


Modules are how Go manages dependencies.

A module is a collection of packages that are released, versioned, and
distributed together. Modules may be downloaded directly from version control
repositories or from module proxy servers.

For a series of tutorials on modules, see
https://golang.org/doc/tutorial/create-module.

For a detailed reference on modules, see https://golang.org/ref/mod.

By default, the go command may download modules from https://proxy.golang.org.
It may authenticate modules using the checksum database at
https://sum.golang.org. Both services are operated by the Go team at Google.
The privacy policies for these services are available at
https://proxy.golang.org/privacy and https://sum.golang.org/privacy,
respectively.

The go command's download behavior may be configured using GOPROXY, GOSUMDB,
GOPRIVATE, and other environment variables. See 'go help environment'
and https://golang.org/ref/mod#private-module-privacy for more information.

*/
