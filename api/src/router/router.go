package router

import (
	"api/src/router/rotas"

	"github.com/gorilla/mux"
)

// SECTION - Gerar vai retornar um router com as rotas configuradas
func Gerar() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)
}
