package poker

import "testing"

type EsbocoArmazenamentoJogador struct {
	pontuações        map[string]int
	chamadasDeVitoria []string
	liga              []Jogador
}

func (s *EsbocoArmazenamentoJogador) ObterLiga() Liga {
	return s.liga
}

func (s *EsbocoArmazenamentoJogador) ObtemPontuacaoDoJogador(nome string) int {
	pontuação := s.pontuações[nome]
	return pontuação
}

func (s *EsbocoArmazenamentoJogador) SalvaVitoria(nome string) {
	s.chamadasDeVitoria = append(s.chamadasDeVitoria, nome)
}

func VerificaVitoriaJogador(t *testing.T, armazenamento *EsbocoArmazenamentoJogador, vencedor string) {
	t.Helper()

	if len(armazenamento.chamadasDeVitoria) != 1 {
		t.Fatalf("recebi %d chamadas de SalvaVitoria esperava %d", len(armazenamento.chamadasDeVitoria), 1)
	}

	if armazenamento.chamadasDeVitoria[0] != vencedor {
		t.Errorf("nao armazenou o vencedor correto, recebi '%s' esperava '%s'", armazenamento.chamadasDeVitoria[0], vencedor)
	}
}
