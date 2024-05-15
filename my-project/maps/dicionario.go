package main

import "errors"

var (
	ErrNaoEncontrado    = errors.New("não foi possível encontrar a palavra")
	ErrPalavraExistente = errors.New("impossível, palavra adionada anteriormente")
)

type Dicionario map[string]string

func (d Dicionario) Busca(palavra string) (string, error) {
	definicao, existe := d[palavra]
	if !existe {
		return "", ErrNaoEncontrado
	}
	return definicao, nil
}

func (d Dicionario) Adiciona(palavra, definicao string) error {
	_, err := d.Busca(palavra)
	switch err {
	case ErrNaoEncontrado:
		d[palavra] = definicao
	case nil:
		return ErrPalavraExistente
	default:
		return err
	}

	return nil
}
