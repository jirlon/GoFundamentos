package poker

import (
	"bufio"
	"io"
	"strings"
)

type CLI struct {
	armazenamentoJogador ArmazenamentoJogador
	in                   io.Reader
}

func (cli *CLI) JogarPoker() {
	reader := bufio.NewScanner(cli.in)
	reader.Scan()
	cli.armazenamentoJogador.SalvaVitoria(extrairVencedor(reader.Text()))
}

func extrairVencedor(userInput string) string {
	return strings.Replace(userInput, " venceu", "", 1)
}
