package main

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    router.Use(cors.Default())


    router.GET("/ping", func(c *gin.Context){
        c.JSON(http.StatusOK, gin.H{
            "message": "pong from docker",
        })
    })
    if err := router.Run(":8080"); err != nil {
        panic(err)
    }
}