package main

import (
	"io"
	"strings"
	"testing"
)

type SistemaDeArquivoDeArmazenamentoDoJogador struct {
	bancoDeDados io.ReadSeeker
}

func (f *SistemaDeArquivoDeArmazenamentoDoJogador) ObterPontuacaoDoJogador(nome string) int {
	var vitorias int

	for _, jogador := range f.PegaLiga() {
		if jogador.Nome == nome {
			vitorias = jogador.Vitorias
			break
		}
	}

	return vitorias
}

func TestSistemaDeArquivoDeArmazenamentoDoJogador(t *testing.T) {

	t.Run("/liga de um leitor", func(t *testing.T) {
		bancoDeDados := strings.NewReader(`[
			{"Nome": "Cleo", "Vitorias": 10},
			{"Nome": "Chris", "Vitorias": 33}
		]`)

		armazenamento := SistemaDeArquivoDeArmazenamentoDoJogador{bancoDeDados}

		recebido := armazenamento.PegaLiga()

		esperado := []Jogador{
			{"Cleo", 10},
			{"Chris", 33},
		}

		defineLiga(t, recebido, esperado)

		//read again
		recebido = armazenamento.PegaLiga()
		defineLiga(t, recebido, esperado)
	})

	t.Run("pegar pontuação do jogador", func(t *testing.T) {
		bancoDeDados := strings.NewReader(`[
			{"Nome": "Cleo", "Vitorias": 10},
			{"Nome": "Chris", "Vitorias": 33}
		]`)

		armazenamento := SistemaDeArquivoDeArmazenamentoDoJogador{bancoDeDados}

		recebido := armazenamento.ObterPontuacaoDoJogador("Chris")

		esperado := 33

		if recebido != esperado {
			t.Errorf("recebido %d esperado %d", recebido, esperado)
		}
	})
}

func (f *SistemaDeArquivoDeArmazenamentoDoJogador) PegaLiga() []Jogador {
	f.bancoDeDados.Seek(0, 0)
	liga, _ := NovaLiga(f.bancoDeDados)
	return liga
}
