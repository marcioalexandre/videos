package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var produtos = []produto{
	{ID: 0, NOME: "teste", ATIVO: false},
}

func hello(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, "Meu primeiro endpoint GET em Golang!")
}

func getProdutos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, produtos)
}

func addProduto(context *gin.Context) {
	var novoProduto produto
	context.BindJSON(&novoProduto)
	produtos = append(produtos, novoProduto)
	context.IndentedJSON(http.StatusCreated, "O produto foi registrado com sucesso!")
}

func main() {
	router := gin.Default()
	router.POST("/produto", addProduto)
	router.GET("/produtos", getProdutos)
	router.GET("/hello", hello)
	router.Run("localhost:8080")
}
