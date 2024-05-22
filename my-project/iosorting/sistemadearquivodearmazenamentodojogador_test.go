package main

import (
	"io"
)

type SistemaDeArquivoDeArmazenamentoDoJogador struct {
	bancoDeDados io.ReadSeeker
}

func (f *SistemaDeArquivoDeArmazenamentoDoJogador) PegaLiga() []Jogador {
	f.bancoDeDados.Seek(0, 0)
	liga, _ := NovaLiga(f.bancoDeDados)
	return liga
}
