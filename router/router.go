package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Inicializa o Router utilizando AS CONFIGURACOES dEFAULT DO GIN
	router := gin.Default()
	// Definindo uma Rota
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// Estamos rodando a nossa api
	router.Run(":8000")
}
