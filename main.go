package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Definir Atributos
type publicacao struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// Definir a variavel Array
var publicacoes = []publicacao{
	{ID: "1", Title: "Blue Train", Content: "John Coltrane"},
	{ID: "2", Title: "Jeru", Content: "Gerry Mulligan"},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Content: "Sarah Vaughan"},
}

// Definir servidor: router (servidor) router.GET (definir rota)
func main() {
	router := gin.Default()
	router.GET("/publicacoes", getPublicacoes)
	router.POST("/publicacoes", postPublicacoes)
	router.Run("localhost:8080")
}

func getPublicacoes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, publicacoes)
}

func postPublicacoes(c *gin.Context) {
	var newPublicacao publicacao
	if err := c.BindJSON(&newPublicacao); err != nil {
		return
	}

	publicacoes = append(publicacoes, newPublicacao)
	c.IndentedJSON(http.StatusCreated, newPublicacao)

}
