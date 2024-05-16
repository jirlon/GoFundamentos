package main

import (
	"bytes"
	"fmt"
	"testing"
)

func TestContagem(t *testing.T) {
	buffer := &bytes.Buffer{}

	Contagem(buffer)

	resultado := buffer.String()
	esperado := "3"

	if resultado != esperado {
		t.Errorf("resultado '%s', esperando '%s'", resultado, esperado)
	}
}

func Contagem(saida *bytes.Buffer) {
	fmt.Fprint(saida, "3")
}
