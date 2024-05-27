package poker

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestSistemaDeArquivoDeArmazenamentoDoJogador(t *testing.T) {

	t.Run("/liga ordenada", func(t *testing.T) {
		bancoDeDados, limpaBancoDeDados := criaArquivoTemporario(t, `[
			{"Nome": "Cleo", "Vitorias": 10},
			{"Nome": "Chris", "Vitorias": 33}
		]`)
		defer limpaBancoDeDados()

		armazenamento, err := NovoSistemaDeArquivoDeArmazenamentoDoJogador(bancoDeDados)

		defineSemErro(t, err)

		recebido := armazenamento.ObterLiga()

		esperado := []Jogador{
			{"Chris", 33},
			{"Cleo", 10},
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

		armazenamento, err := NovoSistemaDeArquivoDeArmazenamentoDoJogador(bancoDeDados)

		defineSemErro(t, err)

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

		armazenamento, err := NovoSistemaDeArquivoDeArmazenamentoDoJogador(bancoDeDados)

		defineSemErro(t, err)

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

		armazenamento, err := NovoSistemaDeArquivoDeArmazenamentoDoJogador(bancoDeDados)

		defineSemErro(t, err)

		armazenamento.SalvaVitoria("Pepper")

		recebido := armazenamento.ObtemPontuacaoDoJogador("Pepper")
		esperado := 1
		definePontuacaoIgual(t, recebido, esperado)
	})

	t.Run("funciona com um arquivo vazio", func(t *testing.T) {
		bancoDeDados, limpaBancoDeDados := criaArquivoTemporario(t, "")
		defer limpaBancoDeDados()

		_, err := NovoSistemaDeArquivoDeArmazenamentoDoJogador(bancoDeDados)

		defineSemErro(t, err)
	})
}

func definePontuacaoIgual(t *testing.T, recebido, esperado int) {
	t.Helper()
	if recebido != esperado {
		t.Errorf("recebido %d esperado %d", recebido, esperado)
	}
}

func criaArquivoTemporario(t *testing.T, dadoInicial string) (*os.File, func()) {
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

func defineSemErro(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("não esperava um erro, mas obteve um, %v", err)
	}
}
