package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kacperkg/vellum/internal/handlers"
)

func Register(router *gin.Engine, authHandler *handlers.AuthHandler) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong from docker2",
		})
	})

	router.POST("/register", authHandler.Register)
	router.POST("/login", authHandler.Login)
}