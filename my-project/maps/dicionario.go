package main

import "errors"

type Dicionario map[string]string

func (d Dicionario) Busca(palavra string) (string, error) {
	definicao, existe := d[palavra]
	if !existe {
		return "", errors.New("não foi posssível encontrar a palavra que voce procura")
	}
	return definicao, nil
}
