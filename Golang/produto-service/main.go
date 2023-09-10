package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func hello(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, "Meu primeiro endpoint GET em Golang!")
}

func main() {
	router := gin.Default()
	router.GET("/hello", hello)
	router.Run("localhost:8080")
}
