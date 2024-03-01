package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	err := httpServer.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
