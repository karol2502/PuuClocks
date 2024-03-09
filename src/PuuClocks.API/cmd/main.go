package main

import (
	"log"
	"net/http"
	"puuclocks/internal/infrastructure"
	"puuclocks/internal/repository"
	"puuclocks/internal/sockets"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

	lobbyManager := sockets.NewLobbyManager()
	r.POST("/create-lobby", func(c *gin.Context) {
		lobby := lobbyManager.CreateLobby()
		c.JSON(http.StatusOK, gin.H{
			"lobbyID": lobby.GetID(),
		})
	})

	r.GET("/join-lobby/:id", func(c *gin.Context) {
		conn, err := sockets.Upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Println(err)
			return
		}
		defer conn.Close()
		id := c.Param("id")

		parsedID, err := uuid.Parse(id)
		if err != nil {
			conn.WriteJSON(map[string]string{
				"message": "User not passed lobby UUID",
			})
		} else {
			l := lobbyManager.FindLobby(parsedID)
			if l == nil {
				conn.WriteJSON(map[string]string{
					"message": "Lobby not found",
				})
			} else {
				sockets.NewClient(conn, l)
				conn.WriteJSON(map[string]string{
					"message": "User connected",
				})
			}
		}
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
