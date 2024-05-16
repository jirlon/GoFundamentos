package main

import (
	"bytes"
	"testing"
)

func TestContagem(t *testing.T) {
	buffer := &bytes.Buffer{}

	Contagem(buffer)

	resultado := buffer.String()
	esperado := `3
2
1
Go!`

	if resultado != esperado {
		t.Errorf("resultado '%s', esperando '%s'", resultado, esperado)
	}
}
