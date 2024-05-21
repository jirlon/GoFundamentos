package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type EsbocoArmazenamentoJogador struct {
	pontuações        map[string]int
	chamadasDeVitoria []string
	liga              []Jogador
}

func (s *EsbocoArmazenamentoJogador) ObterLiga() []Jogador {
	return s.liga
}

func (s *EsbocoArmazenamentoJogador) ObtemPontuacaoDoJogador(nome string) int {
	pontuação := s.pontuações[nome]
	return pontuação
}

func (s *EsbocoArmazenamentoJogador) GravarVitoria(nome string) {
	s.chamadasDeVitoria = append(s.chamadasDeVitoria, nome)
}

func TestObterJogadores(t *testing.T) {
	armazenamento := EsbocoArmazenamentoJogador{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		nil,
		nil,
	}
	servidor := NovoServidorJogador(&armazenamento)

	t.Run("retorna pontuação de Pepper", func(t *testing.T) {
		requisicao := novaRequisicaoObterPontuacao("Pepper")
		resposta := httptest.NewRecorder()

		servidor.ServeHTTP(resposta, requisicao)

		verificaStatus(t, resposta.Code, http.StatusOK)
		verificaCorpoDaResposta(t, resposta.Body.String(), "20")
	})

	t.Run("retorna pontuação do Floyd", func(t *testing.T) {
		requisicao := novaRequisicaoObterPontuacao("Floyd")
		resposta := httptest.NewRecorder()

		servidor.ServeHTTP(resposta, requisicao)

		verificaStatus(t, resposta.Code, http.StatusOK)
		verificaCorpoDaResposta(t, resposta.Body.String(), "10")
	})

	t.Run("retorna 404 para jogadores em falta", func(t *testing.T) {
		requisicao := novaRequisicaoObterPontuacao("Apollo")
		resposta := httptest.NewRecorder()

		servidor.ServeHTTP(resposta, requisicao)

		verificaStatus(t, resposta.Code, http.StatusNotFound)
	})
}

func TestArmazenarVitórias(t *testing.T) {
	armazenamento := EsbocoArmazenamentoJogador{
		map[string]int{},
		nil,
		nil,
	}
	servidor := NovoServidorJogador(&armazenamento)

	t.Run("grava vitória no POST", func(t *testing.T) {
		jogador := "Pepper"

		requisicao := novaRequisiçãoPostDeVitoria(jogador)
		resposta := httptest.NewRecorder()

		servidor.ServeHTTP(resposta, requisicao)

		verificaStatus(t, resposta.Code, http.StatusAccepted)

		if len(armazenamento.chamadasDeVitoria) != 1 {
			t.Fatalf("obteve %d chamadas para GravarVitoria, esperava %d", len(armazenamento.chamadasDeVitoria), 1)
		}

		if armazenamento.chamadasDeVitoria[0] != jogador {
			t.Errorf("não armazenou o vencedor correto, obteve '%s', esperava '%s'", armazenamento.chamadasDeVitoria[0], jogador)
		}
	})
}

func TestLiga(t *testing.T) {
	t.Run("retorna a tabela da liga como json", func(t *testing.T) {
		ligaEsperada := []Jogador{
			{"Cleo", 32},
			{"Chris", 20},
			{"Tiest", 14},
		}

		armazenamento := EsbocoArmazenamentoJogador{nil, nil, ligaEsperada}
		servidor := NovoServidorJogador(&armazenamento)

		requisicao, _ := http.NewRequest(http.MethodGet, "/liga", nil)
		resposta := httptest.NewRecorder()

		servidor.ServeHTTP(resposta, requisicao)

		var obtido []Jogador

		err := json.NewDecoder(resposta.Body).Decode(&obtido)

		if err != nil {
			t.Fatalf("nao foi possivel fazer parse da resposta do servidor '%s' no slice de Jogador, '%v'", resposta.Body, err)
		}

		verificaStatus(t, resposta.Code, http.StatusOK)

		if !reflect.DeepEqual(obtido, ligaEsperada) {
			t.Errorf("obtido %v esperado %v", obtido, ligaEsperada)
		}
	})
}

func verificaStatus(t *testing.T, obtido, esperado int) {
	t.Helper()
	if obtido != esperado {
		t.Errorf("não obteve o status correto, obteve %d, esperava %d", obtido, esperado)
	}
}

func novaRequisicaoObterPontuacao(nome string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/jogadores/%s", nome), nil)
	return req
}

func novaRequisiçãoPostDeVitoria(nome string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/jogadores/%s", nome), nil)
	return req
}

func verificaCorpoDaResposta(t *testing.T, obtido, esperado string) {
	t.Helper()
	if obtido != esperado {
		t.Errorf("resposta corpo está incorreta, obtido '%s' esperado '%s'", obtido, esperado)
	}
}
