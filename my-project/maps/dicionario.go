package main

import "errors"

var ErrNaoEncontrado = errors.New("não foi possível encontrar a palavra")

type Dicionario map[string]string

func (d Dicionario) Busca(palavra string) (string, error) {
	definicao, existe := d[palavra]
	if !existe {
		return "", ErrNaoEncontrado
	}
	return definicao, nil
}

func (d Dicionario) Adiciona(palavra, definicao string) {
	d[palavra] = definicao
}
