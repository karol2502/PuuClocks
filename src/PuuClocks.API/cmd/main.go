package main

import (
	"net/http"
	"puuclocks/internal/infrastructure"
	"puuclocks/internal/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	dbCfg := repository.DatabasesConfig{
		RedisConfig: repository.RedisConfig{
			Addr: "redis:6379",
		},
		MySqlConfig: infrastructure.MySqlConfig{
			DBName: "mysql",
			Path: "root:root@tcp(mysql:3306)/puuclocks",
		},
	}

	_, err := repository.NewDatabases(dbCfg)
	if err != nil {
		panic(err)
	}

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

	err = httpServer.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
