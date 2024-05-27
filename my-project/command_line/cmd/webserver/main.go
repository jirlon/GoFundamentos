package main

import (
	"fmt"
	"log"
	"os"

	poker "github.com/jirlon/GoFundamentos/command_line"
)

const dbFileName = "game.db.json"

func main() {
	fmt.Println("Vamos jogar poker")
	fmt.Println("Digite {Nome} venceu para registrar uma vitoria")

	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problema abrindo %s %v", dbFileName, err)
	}
	armazenamento, err := poker.NovoSistemaDeArquivoDeArmazenamentoDoJogador(db)

	if err != nil {
		log.Fatalf("problema criando o sistema de arquivo do armazenamento do jogador, %v", err)
	}

	jogo := poker.NovoCLI(armazenamento, os.Stdin)
	jogo.JogarPoker()
	/*
		servidor := poker.NovoServidorJogador(armazenamento)

		if err := http.ListenAndServe(":5000", servidor); err != nil {
			log.Fatalf("não foi possível ouvir na porta 5000 %v", err)
		}*/
}
