package main

import (
	"fmt"
	"log"
	"os"

	poker "github.com/jirlon/GoFundamentos/command_line"
)

const dbFileName = "game.db.json"

func main() {

	armazenamento, close, err := poker.ArmazenamentoSistemaDeArquivoJogadorAPartirDeArquivo(dbFileName)

	if err != nil {
		log.Fatal(err)
	}
	defer close()

	fmt.Println("Vamos jogar poker")
	fmt.Println("Digite {Nome} venceu para registrar uma vitoria")

	poker.NovoCLI(armazenamento, os.Stdin).JogarPoker()
}
