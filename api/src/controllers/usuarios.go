package controllers

import "net/http"

// SECTION - Insere um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Criando usuário!"))
}

// SECTION - Busca os usuários no banco de dados
func BuscarUsuarios(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Buscando todos os usuários!"))
}

// SECTION - Busca um usuário no banco de dados
func BuscarUsuario(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Buscando um usuário!"))
}

// SECTION - Atualiza as informações de um usuário no banco de dados
func AtualizarUsuario(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Atualizando usuário!"))
}

// SECTION - Deleta um usuário no banco de dados
func DeletarUsuario(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Deletando usuário!"))
}