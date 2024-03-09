package main

import (
	"net/http"
	"puuclocks/internal/infrastructure"
	"puuclocks/internal/log"
	"puuclocks/internal/repository"
	"puuclocks/internal/sockets"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

func main() {
	dbCfg := repository.DatabasesConfig{
		RedisConfig: repository.RedisConfig{
			Addr: "redis:6379",
		},
		MySQLConfig: infrastructure.MySQLConfig{
			DBName: "mysql",
			Path:   "root:root@tcp(mysql:3306)/puuclocks",
		},
	}

	log.InitLogger()

	_, err := repository.NewDatabases(&dbCfg)
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
		var conn *websocket.Conn
		var parsedID uuid.UUID
		conn, err = sockets.Upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Log.Errorln(err)
			return
		}
		defer conn.Close()
		id := c.Param("id")

		parsedID, err = uuid.Parse(id)
		if err != nil {
			if err = conn.WriteJSON(map[string]string{
				"message": "User not passed lobby UUID",
			}); err != nil {
				log.Log.Errorln(err)
			}
		} else {
			l := lobbyManager.FindLobby(parsedID)
			if l == nil {
				if err = conn.WriteJSON(map[string]string{
					"message": "Lobby not found",
				}); err != nil {
					log.Log.Errorln(err)
				}
			} else {
				sockets.NewClient(conn, l)
				if err = conn.WriteJSON(map[string]string{
					"message": "User connected",
				}); err != nil {
					log.Log.Errorln(err)
				}
			}
		}
	})

	httpServer := &http.Server{
		Addr:              ":8080",
		Handler:           r,
		ReadHeaderTimeout: time.Second,
	}

	err = httpServer.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
