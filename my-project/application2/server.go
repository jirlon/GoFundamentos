package main

import (
	"fmt"
	"net/http"
)

type ArmazenamentoJogador interface {
	ObtemPontuacaoDoJogador(nome string) int
	GravarVitoria(nome string)
}

type ServidorJogador struct {
	armazenamento ArmazenamentoJogador
}

func (s *ServidorJogador) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	roteador := http.NewServeMux()

	roteador.Handle("/liga", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	roteador.Handle("/jogadores/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		jogador := r.URL.Path[len("/jogadores/"):]

		switch r.Method {
		case http.MethodPost:
			s.processarVitoria(w, jogador)
		case http.MethodGet:
			s.mostrarPontuacao(w, jogador)
		}
	}))

	roteador.ServeHTTP(w, r)
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
