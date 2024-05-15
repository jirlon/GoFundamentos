package maps

import "testing"

func Busca(dicionario map[string]string, palavra string) string {
	return dicionario[palavra]
}

func TestBusca(t *testing.T) {
	dicionario := map[string]string{"teste": "isso é apenas um teste"}

	resultado := Busca(dicionario, "teste")
	esperado := "isso é apenas um teste"

	if resultado != esperado {
		t.Errorf("resultado '%s', esperado '%s', dado '%s'", resultado, esperado, "test")
	}
}
