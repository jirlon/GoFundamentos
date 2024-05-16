package main

import (
	"fmt"
	"io"
	"os"
)

func Contagem(saida io.Writer) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(saida, i)
	}
	fmt.Fprint(saida, "Go!")
}

func main() {
	Contagem(os.Stdout)
}
