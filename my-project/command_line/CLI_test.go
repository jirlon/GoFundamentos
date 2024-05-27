package poker

import (
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	in := strings.NewReader("Chris venceu\n")
	ArmazenamentoJogador := &EsbocoArmazenamentoJogador{}

	cli := &CLI{ArmazenamentoJogador, in}
	cli.JogarPoker()

	if len(ArmazenamentoJogador.chamadasDeVitoria) < 1 {
		t.Fatal("esperando uma chamada de vitoria mas nao recebi nenhuma")
	}

	obtido := ArmazenamentoJogador.chamadasDeVitoria[0]
	esperado := "Chris"

	if obtido != esperado {
		t.Errorf("nao armazenou o vencedor correto, recebi '%s', esperava '%s'", obtido, esperado)
	}
}
