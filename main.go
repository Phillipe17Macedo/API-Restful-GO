package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type Pessoa struct {
	ID   int    `json:"id"`
	Nome string `json:"nome"`
}

var pessoas []Pessoa

func getListPessoas(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(pessoas)
}

func getPessoa(w http.ResponseWriter, r *http.Request) {
	nome := r.URL.Query().Get("nome")
	idStr := r.URL.Query().Get("id")

	if idStr != "" {
		id, _ := strconv.Atoi(idStr)

		for _, pessoa := range pessoas {
			if pessoa.ID == id {
				json.NewEncoder(w).Encode(pessoa)
				return
			}
		}
		json.NewEncoder(w).Encode(nil)
		return
	}

	if nome != "" {
		for _, pessoa := range pessoas {
			if pessoa.Nome == nome {
				json.NewEncoder(w).Encode(pessoa)
				return
			}
		}
		json.NewEncoder(w).Encode(nil)
		return
	}
}

func postPessoa(w http.ResponseWriter, r *http.Request) {
	var pessoa Pessoa
	_ = json.NewDecoder(r.Body).Decode(&pessoa)
	pessoa.ID = len(pessoas) + 1
	pessoas = append(pessoas, pessoa)
	json.NewEncoder(w).Encode(pessoa)
}

func deletePessoa(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.URL.Query().Get("id"))
	var novaSlice []Pessoa
	for _, pessoa := range pessoas {
		if pessoa.ID != id {
			novaSlice = append(novaSlice, pessoa)
		}
	}

	for i, pessoa := range novaSlice {
		pessoa.ID = i + 1
		novaSlice[i] = pessoa
	}

	pessoas = novaSlice
	fmt.Fprintf(w, "Pessoa com ID %d foi deletada com sucesso.", id)
}

func main() {
	http.HandleFunc("/getListPessoas", getListPessoas)
	http.HandleFunc("/getPessoa", getPessoa)
	http.HandleFunc("/postPessoa", postPessoa)
	http.HandleFunc("/deletePessoa", deletePessoa)

	pessoas = append(pessoas, Pessoa{ID: 1, Nome: "Phillipe"})
	pessoas = append(pessoas, Pessoa{ID: 2, Nome: "Camargos"})
	pessoas = append(pessoas, Pessoa{ID: 3, Nome: "João VItor"})
	pessoas = append(pessoas, Pessoa{ID: 4, Nome: "Gabriela"})
	pessoas = append(pessoas, Pessoa{ID: 5, Nome: "Rocha"})
	pessoas = append(pessoas, Pessoa{ID: 6, Nome: "Jean"})

	_ = http.ListenAndServe(":3333", nil)
}