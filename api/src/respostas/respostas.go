package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

// ANCHOR - Retorna uma resposta em JSON para a requisição
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if erro := json.NewEncoder(w).Encode(dados); erro != nil {
		log.Fatal(erro)
	}
}

// ANCHOR - Retorna uma resposta de erro em JSON para a requisição
func Erro(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode, struct{
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(),
	})
}
