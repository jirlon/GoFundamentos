package main

const (
	ErrNaoEncontrado      = ErrDicionario("não foi possível encontrar a palavra")
	ErrPalavraExistente   = ErrDicionario("impossível, palavra adionada anteriormente")
	ErrPalavraInexistente = ErrDicionario("impossível, palavra não existe")
)

type Dicionario map[string]string

type ErrDicionario string

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

func (e ErrDicionario) Error() string {
	return string(e)
}

func (d Dicionario) Atualiza(palavra, definicao string) error {
	_, err := d.Busca(palavra)
	switch err {
	case ErrNaoEncontrado:
		return ErrPalavraInexistente
	case nil:
		d[palavra] = definicao
	default:
		return err
	}

	return nil
}

func (d Dicionario) Deleta(palavra string) {
	delete(d, palavra)
}
