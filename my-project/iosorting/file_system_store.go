package main

import (
	"encoding/json"
	"os"
)

type SistemaDeArquivoDeArmazenamentoDoJogador struct {
	bancoDeDados *json.Encoder
	liga         Liga
}

func NovoSistemaDeArquivoDeArmazenamentoDoJogador(bancoDeDados *os.File) *SistemaDeArquivoDeArmazenamentoDoJogador {
	bancoDeDados.Seek(0, 0)
	liga, _ := NovaLiga(bancoDeDados)
	return &SistemaDeArquivoDeArmazenamentoDoJogador{
		bancoDeDados: json.NewEncoder(&fita{bancoDeDados}),
		liga:         liga,
	}
}

func (f *SistemaDeArquivoDeArmazenamentoDoJogador) ObtemPontuacaoDoJogador(nome string) int {

	jogador := f.liga.Find(nome)

	if jogador != nil {
		return jogador.Vitorias
	}
	return 0
}

func (f *SistemaDeArquivoDeArmazenamentoDoJogador) ObterLiga() Liga {
	return f.liga
}

func (f *SistemaDeArquivoDeArmazenamentoDoJogador) SalvaVitoria(nome string) {
	jogador := f.liga.Find(nome)

	if jogador != nil {
		jogador.Vitorias++
	} else {
		f.liga = append(f.liga, Jogador{nome, 1})
	}

	//f.bancoDeDados.Seek(0, 0)
	//json.NewEncoder(f.bancoDeDados).Encode(f.liga)
	f.bancoDeDados.Encode(f.liga)
}
