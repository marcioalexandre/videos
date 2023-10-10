package main

import (
	"fmt"
	"net/http"
	"strconv"

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

func parseToInt(id string) int {
	id64, err := strconv.ParseInt(id, 10, 8)
	if err != nil {
		fmt.Println(err)
	}
	return int(id64)
}

func getIndexById(id int) int {
	index := -1
	for i, produto := range produtos {
		if produto.ID == id {
			index = i
		}
	}
	return index
}

func getProduto(context *gin.Context) {
	id := context.Param("id") // formato string
	index := getIndexById(parseToInt(id))

	if index > -1 {
		context.IndentedJSON(http.StatusFound, produtos[index])
		return
	}

	context.IndentedJSON(http.StatusNotFound, "Produto inexistente!")
}

func deleteProduto(context *gin.Context) {
	id := context.Param("id")
	index := getIndexById(parseToInt(id))

	if index > -1 {
		produtos = append(produtos[:index], produtos[index+1:]...)
		context.IndentedJSON(http.StatusOK, "Produto excluído com sucesso!")
		return
	}

	context.IndentedJSON(http.StatusNotFound, "Produto inexistente para exclusão!")
}

func main() {
	router := gin.Default()
	router.POST("/produto", addProduto)
	router.GET("/produto", getProdutos)
	router.GET("/hello", hello)
	router.GET("/produto/:id", getProduto)
	router.DELETE("/produto/:id", deleteProduto)
	router.Run("localhost:8080")
}
