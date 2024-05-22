package main

import (
	"encoding/json"
	"io"
)

type SistemaDeArquivoDeArmazenamentoDoJogador struct {
	bancoDeDados io.Reader
}

func (f *SistemaDeArquivoDeArmazenamentoDoJogador) PegaLiga() []Jogador {
	var liga []Jogador
	json.NewDecoder(f.bancoDeDados).Decode(&liga)
	return liga
}
