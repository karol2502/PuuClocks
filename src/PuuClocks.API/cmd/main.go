package main

import (
	"fmt"
	"net/http"
	"puuclocks/internal/infrastructure"
	"puuclocks/internal/repository"
	"puuclocks/internal/sockets"

	"github.com/gin-gonic/gin"
)

func main() {
	dbCfg := repository.DatabasesConfig{
		RedisConfig: repository.RedisConfig{
			Addr: "redis:6379",
		},
		MySqlConfig: infrastructure.MySqlConfig{
			DBName: "mysql",
			Path:   "root:root@tcp(mysql:3306)/puuclocks",
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

	lobby := sockets.NewLobby()
	r.GET("/ws", func(c *gin.Context) {
		fmt.Println(len(lobby.Clients) + 1, " Users")
		lobby.JoinLobby(c.Writer, c.Request)
	})
	go lobby.Run()

	httpServer := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	err = httpServer.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
