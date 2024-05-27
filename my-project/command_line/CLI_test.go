package poker

import (
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {

	t.Run("recorda vencedor chris digitado pelo usuario", func(t *testing.T) {
		in := strings.NewReader("Chris venceu\n")
		armazenamentoJogador := &EsbocoArmazenamentoJogador{}

		cli := &CLI{armazenamentoJogador, in}
		cli.JogarPoker()

		verificaVitoriaJogador(t, armazenamentoJogador, "Chris")
	})

	t.Run("recorda vencedor cleo digitado pelo usuario", func(t *testing.T) {
		in := strings.NewReader("Cleo venceu\n")
		armazenamentoJogador := &EsbocoArmazenamentoJogador{}

		cli := &CLI{armazenamentoJogador, in}
		cli.JogarPoker()

		verificaVitoriaJogador(t, armazenamentoJogador, "Cleo")
	})
}
