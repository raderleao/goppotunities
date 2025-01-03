package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// Initialize Router
	router := gin.Default()

	// Definindo uma Rota
	InitializeRoutes(router)

	// Run the server
	router.Run(":8000")
}
