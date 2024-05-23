package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"testing"
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

func TestSistemaDeArquivoDeArmazenamentoDoJogador(t *testing.T) {

	t.Run("/liga de um leitor", func(t *testing.T) {
		bancoDeDados, limpaBancoDeDados := criaArquivoTemporario(t, `[
			{"Nome": "Cleo", "Vitorias": 10},
			{"Nome": "Chris", "Vitorias": 33}
		]`)
		defer limpaBancoDeDados()

		armazenamento := SistemaDeArquivoDeArmazenamentoDoJogador{bancoDeDados}

		recebido := armazenamento.ObterLiga()

		esperado := []Jogador{
			{"Cleo", 10},
			{"Chris", 33},
		}

		defineLiga(t, recebido, esperado)

		//read again
		recebido = armazenamento.ObterLiga()
		defineLiga(t, recebido, esperado)
	})

	t.Run("retorna pontuação do jogador", func(t *testing.T) {
		bancoDeDados, limpaBancoDeDados := criaArquivoTemporario(t, `[
			{"Nome": "Cleo", "Vitorias": 10},
			{"Nome": "Chris", "Vitorias": 33}
		]`)
		defer limpaBancoDeDados()

		armazenamento := SistemaDeArquivoDeArmazenamentoDoJogador{bancoDeDados}

		recebido := armazenamento.ObtemPontuacaoDoJogador("Chris")

		esperado := 33

		definePontuacaoIgual(t, recebido, esperado)
	})

	t.Run("armazena vitorias de um jogador existente", func(t *testing.T) {
		bancoDeDados, limpaBancoDeDados := criaArquivoTemporario(t, `[
			{"Nome": "Cleo", "Vitorias": 10},
			{"Nome": "Chris", "Vitorias": 33}
		]`)
		defer limpaBancoDeDados()

		armazenamento := SistemaDeArquivoDeArmazenamentoDoJogador{bancoDeDados}

		armazenamento.SalvaVitoria("Chris")

		recebido := armazenamento.ObtemPontuacaoDoJogador("Chris")
		esperado := 34
		definePontuacaoIgual(t, recebido, esperado)
	})

	t.Run("armazena vitorias de novos jogadores", func(t *testing.T) {
		bancoDeDados, limpaBancoDeDados := criaArquivoTemporario(t, `[
			{"Nome": "Cleo", "Vitorias": 10},
			{"Nome": "Chris", "Vitorias": 33}
		]`)
		defer limpaBancoDeDados()

		armazenamento := SistemaDeArquivoDeArmazenamentoDoJogador{bancoDeDados}

		armazenamento.SalvaVitoria("Pepper")

		recebido := armazenamento.ObtemPontuacaoDoJogador("Pepper")
		esperado := 1
		definePontuacaoIgual(t, recebido, esperado)
	})
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

func definePontuacaoIgual(t *testing.T, recebido, esperado int) {
	t.Helper()
	if recebido != esperado {
		t.Errorf("recebido %d esperado %d", recebido, esperado)
	}
}

func criaArquivoTemporario(t *testing.T, dadoInicial string) (io.ReadWriteSeeker, func()) {
	t.Helper()

	arquivotmp, err := ioutil.TempFile("", "db")

	if err != nil {
		t.Fatalf("não foi possivel escrever o arquivo temporario %v", err)
	}

	arquivotmp.Write([]byte(dadoInicial))

	removeArquivo := func() {
		arquivotmp.Close()
		os.Remove(arquivotmp.Name())
	}

	return arquivotmp, removeArquivo
}

func (f *SistemaDeArquivoDeArmazenamentoDoJogador) ObterLiga() Liga {
	f.bancoDeDados.Seek(0, 0)
	liga, _ := NovaLiga(f.bancoDeDados)
	return liga
}
