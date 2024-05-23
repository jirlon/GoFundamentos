package main

func NovoArmazenamentoDeJogadorNaMemoria() *ArmazenamentoDeJogadorNaMemoria {
	return &ArmazenamentoDeJogadorNaMemoria{map[string]int{}}
}

type ArmazenamentoDeJogadorNaMemoria struct {
	armazenamento map[string]int
}

func (a *ArmazenamentoDeJogadorNaMemoria) ObterLiga() Liga {
	var liga []Jogador
	for nome, vitorias := range a.armazenamento {
		liga = append(liga, Jogador{nome, vitorias})
	}

	return liga
}

func (a *ArmazenamentoDeJogadorNaMemoria) SalvaVitoria(nome string) {
	a.armazenamento[nome]++
}

func (a *ArmazenamentoDeJogadorNaMemoria) ObtemPontuacaoDoJogador(nome string) int {
	return a.armazenamento[nome]
}
