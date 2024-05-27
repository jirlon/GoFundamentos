package poker

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

const tipoDoConteudoJSON = "application/json"

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
		verificaVitoriaJogador(t, &armazenamento, jogador)
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

		obtido := obterLigaDaResposta(t, resposta.Body)
		verificaStatus(t, resposta.Code, http.StatusOK)
		verificaLiga(t, obtido, ligaEsperada)

		verificaTipoDoConteudo(t, resposta, tipoDoConteudoJSON)
	})
}

func verificaTipoDoConteudo(t *testing.T, resposta *httptest.ResponseRecorder, esperado string) {
	t.Helper()
	if resposta.Result().Header.Get("content-type") != "application/json" {
		t.Errorf("resposta não tinha o tipo de conteúdo de application/json, obtido %v", resposta.Result().Header)
	}
}

func obterLigaDaResposta(t *testing.T, body io.Reader) (liga []Jogador) {
	t.Helper()
	liga, err := NovaLiga(body)

	if err != nil {
		t.Fatalf("nao foi possivel fazer parse da resposta do servidor '%s' no slice de Jogador, '%v'", body, err)
	}

	return liga
}

func verificaLiga(t *testing.T, obtido, esperado []Jogador) {
	t.Helper()
	if !reflect.DeepEqual(obtido, esperado) {
		t.Errorf("obtido %v esperado %v", obtido, esperado)
	}
}

func verificaStatus(t *testing.T, obtido, esperado int) {
	t.Helper()
	if obtido != esperado {
		t.Errorf("não obteve o status correto, obteve %d, esperava %d", obtido, esperado)
	}
}

func verificaCorpoDaResposta(t *testing.T, obtido, esperado string) {
	t.Helper()
	if obtido != esperado {
		t.Errorf("resposta corpo está incorreta, obtido '%s' esperado '%s'", obtido, esperado)
	}
}

func defineLiga(t *testing.T, recebido, esperado []Jogador) {
	t.Helper()
	if !reflect.DeepEqual(recebido, esperado) {
		t.Errorf("recebido %v esperado %v", recebido, esperado)
	}
}

func novaRequisicaoDeLiga() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/liga", nil)
	return req
}

func novaRequisicaoObterPontuacao(nome string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/jogadores/%s", nome), nil)
	return req
}

func novaRequisiçãoPostDeVitoria(nome string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/jogadores/%s", nome), nil)
	return req
}

func verificaVitoriaJogador(t *testing.T, armazenamento *EsbocoArmazenamentoJogador, vencedor string) {
	t.Helper()

	if len(armazenamento.chamadasDeVitoria) != 1 {
		t.Fatalf("recebi %d chamadas de SalvaVitoria esperava %d", len(armazenamento.chamadasDeVitoria), 1)
	}

	if armazenamento.chamadasDeVitoria[0] != vencedor {
		t.Errorf("nao armazenou o vencedor correto, recebi '%s' esperava '%s'", armazenamento.chamadasDeVitoria[0], vencedor)
	}
}
