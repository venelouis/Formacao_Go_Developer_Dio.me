package main

import (
	"encoding/json" // Importando a biblioteca para manipular JSON
	"fmt"           // Importando a biblioteca para imprimir na tela
	"io/ioutil"     // Importando a biblioteca para ler arquivos
	"os"            // Importando a biblioteca para manipular arquivos
	"strconv"       // Importando a biblioteca para converter tipos de dados
)

type Users struct { // Criando uma estrutura para armazenar os usuários
	Users []User `json:"usuarios"` // Criando um array de usuários
}

type User struct { // Criando uma estrutura para armazenar os dados do usuário
	Nome   string `json:"Nome"`        // Criando um campo para armazenar o nome do usuário
	Tipo   string `json:"Tipo"`        // Criando um campo para armazenar o tipo do usuário
	Idade  int    `json:"Idade"`       // Criando um campo para armazenar a idade do usuário
	social Social `json:"rede social"` // Criando um campo para armazenar a rede social do usuário
}

type Social struct { // Criando uma estrutura para armazenar os dados da rede social do usuário
	facebook string `json:"facebook"` // Criando um campo para armazenar o facebook do usuário
	twitter  string `json:"twitter"`  // Criando um campo para armazenar o twitter do usuário
}

func main() { // Função principal

	jsonFile, err := os.Open("usuarios.json") // Abrindo o arquivo usuarios.json

	if err != nil { // Verificando se houve erro ao abrir o arquivo
		fmt.Println(err) // Imprimindo o erro
	}

	fmt.Println("Arquivo aberto com sucesso") // Imprimindo que o arquivo foi aberto com sucesso

	defer jsonFile.Close() // Fechando o arquivo após a execução da função

	byteValue, _ := ioutil.ReadAll(jsonFile) // Lendo o arquivo

	var usuarios Users // Criando uma variável para armazenar os usuários

	json.Unmarshal(byteValue, &usuarios) // Convertendo o JSON para a estrutura criada

	for i := 0; i < len(usuarios.Users); i++ { // Percorrendo o array de usuários
		fmt.Println("Usuario Tipo: " + usuarios.Users[i].Tipo)                 // Imprimindo o tipo do usuário
		fmt.Println("Usuário Idade: " + strconv.Itoa(usuarios.Users[i].Idade)) // Imprimindo a idade do usuário
		fmt.Println("Usuário Nome: " + usuarios.Users[i].Nome)                 // Imprimindo o nome do usuário

	}

}
