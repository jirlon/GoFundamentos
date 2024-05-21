package main

func NovoArmazenamentoDeJogadorNaMemoria() *ArmazenamentoDeJogadorNaMemoria {
	return &ArmazenamentoDeJogadorNaMemoria{map[string]int{}}
}

type ArmazenamentoDeJogadorNaMemoria struct {
	armazenamento map[string]int
}

func (a *ArmazenamentoDeJogadorNaMemoria) ObterLiga() []Jogador {
	return nil
}

func (a *ArmazenamentoDeJogadorNaMemoria) GravarVitoria(nome string) {
	a.armazenamento[nome]++
}

func (a *ArmazenamentoDeJogadorNaMemoria) ObtemPontuacaoDoJogador(nome string) int {
	return a.armazenamento[nome]
}
