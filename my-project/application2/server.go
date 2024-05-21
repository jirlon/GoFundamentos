package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ArmazenamentoJogador interface {
	ObtemPontuacaoDoJogador(nome string) int
	GravarVitoria(nome string)
}

type ServidorJogador struct {
	armazenamento ArmazenamentoJogador
	http.Handler
}

type Jogador struct {
	Nome     string
	Vitorias int
}

func NovoServidorJogador(armazenamento ArmazenamentoJogador) *ServidorJogador {
	s := new(ServidorJogador)

	s.armazenamento = armazenamento

	roteador := http.NewServeMux()
	roteador.Handle("/liga", http.HandlerFunc(s.manipulaLiga))
	roteador.Handle("/jogadores/", http.HandlerFunc(s.manipulaJogadores))

	s.Handler = roteador

	return s
}

func (s *ServidorJogador) manipulaLiga(w http.ResponseWriter, r *http.Request) {
	tabelaDaLiga := []Jogador{
		{"Chris", 20},
	}

	json.NewEncoder(w).Encode(tabelaDaLiga)

	w.WriteHeader(http.StatusOK)
}

func (s *ServidorJogador) manipulaJogadores(w http.ResponseWriter, r *http.Request) {
	jogador := r.URL.Path[len("/jogadores/"):]

	switch r.Method {
	case http.MethodPost:
		s.processarVitoria(w, jogador)
	case http.MethodGet:
		s.mostrarPontuacao(w, jogador)
	}
}

func (s *ServidorJogador) mostrarPontuacao(w http.ResponseWriter, jogador string) {
	pontuação := s.armazenamento.ObtemPontuacaoDoJogador(jogador)

	if pontuação == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, pontuação)
}

func (s *ServidorJogador) processarVitoria(w http.ResponseWriter, jogador string) {
	s.armazenamento.GravarVitoria(jogador)
	w.WriteHeader(http.StatusAccepted)
}
