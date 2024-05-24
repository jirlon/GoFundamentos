package main

import (
	"encoding/json"
	"io"
)

type SistemaDeArquivoDeArmazenamentoDoJogador struct {
	bancoDeDados io.ReadWriteSeeker
}

func (f *SistemaDeArquivoDeArmazenamentoDoJogador) ObtemPontuacaoDoJogador(nome string) int {

	jogador := f.ObterLiga().Find(nome)

	if jogador != nil {
		return jogador.Vitorias
	}
	return 0
}

func (f *SistemaDeArquivoDeArmazenamentoDoJogador) ObterLiga() Liga {
	f.bancoDeDados.Seek(0, 0)
	liga, _ := NovaLiga(f.bancoDeDados)
	return liga
}

func (f *SistemaDeArquivoDeArmazenamentoDoJogador) SalvaVitoria(nome string) {
	liga := f.ObterLiga()
	jogador := liga.Find(nome)

	if jogador != nil {
		jogador.Vitorias++
	} else {
		liga = append(liga, Jogador{nome, 1})
	}

	f.bancoDeDados.Seek(0, 0)
	json.NewEncoder(f.bancoDeDados).Encode(liga)
}
