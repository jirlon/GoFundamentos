package poker

import "io"

type CLI struct {
	armazenamentoJogador ArmazenamentoJogador
	in                   io.Reader
}

func (cli *CLI) JogarPoker() {
	cli.armazenamentoJogador.SalvaVitoria("Chris")
}
