package poker

import (
	"bufio"
	"io"
	"strings"
)

type CLI struct {
	armazenamentoJogador ArmazenamentoJogador
	in                   *bufio.Scanner
}

func NovoCLI(armazenamento ArmazenamentoJogador, in io.Reader) *CLI {
	return &CLI{
		armazenamentoJogador: armazenamento,
		in:                   bufio.NewScanner(in),
	}
}

func (cli *CLI) JogarPoker() {
	userInput := cli.readLine()
	cli.armazenamentoJogador.SalvaVitoria(extrairVencedor(userInput))
}

func extrairVencedor(userInput string) string {
	return strings.Replace(userInput, " venceu", "", 1)
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}
