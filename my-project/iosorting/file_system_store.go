package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

type SistemaDeArquivoDeArmazenamentoDoJogador struct {
	bancoDeDados *json.Encoder
	liga         Liga
}

// construtor
func NovoSistemaDeArquivoDeArmazenamentoDoJogador(arquivo *os.File) (*SistemaDeArquivoDeArmazenamentoDoJogador, error) {

	err := iniciaArquivoBDDeJogador(arquivo)

	if err != nil {
		return nil, fmt.Errorf("problema inicializando arquivo do jogador, %v", err)
	}

	liga, err := NovaLiga(arquivo)

	if err != nil {
		return nil, fmt.Errorf("problema carregando o armazenamento do jogador de arquivo %s, %v", arquivo.Name(), err)
	}
	return &SistemaDeArquivoDeArmazenamentoDoJogador{
		bancoDeDados: json.NewEncoder(&fita{arquivo}),
		liga:         liga,
	}, nil
}

func iniciaArquivoBDDeJogador(arquivo *os.File) error {
	arquivo.Seek(0, 0)

	info, err := arquivo.Stat()

	if err != nil {
		return fmt.Errorf("problema ao usar o arquivo %s, %v", arquivo.Name(), err)
	}

	if info.Size() == 0 {
		arquivo.Write([]byte("[]"))
		arquivo.Seek(0, 0)
	}

	return nil
}

func (f *SistemaDeArquivoDeArmazenamentoDoJogador) ObtemPontuacaoDoJogador(nome string) int {

	jogador := f.liga.Find(nome)

	if jogador != nil {
		return jogador.Vitorias
	}
	return 0
}

func (f *SistemaDeArquivoDeArmazenamentoDoJogador) ObterLiga() Liga {
	sort.Slice(f.liga, func(i, j int) bool {
		return f.liga[i].Vitorias > f.liga[j].Vitorias
	})
	return f.liga
}

func (f *SistemaDeArquivoDeArmazenamentoDoJogador) SalvaVitoria(nome string) {
	jogador := f.liga.Find(nome)

	if jogador != nil {
		jogador.Vitorias++
	} else {
		f.liga = append(f.liga, Jogador{nome, 1})
	}

	//f.bancoDeDados.Seek(0, 0)
	//json.NewEncoder(f.bancoDeDados).Encode(f.liga)
	f.bancoDeDados.Encode(f.liga)
}
