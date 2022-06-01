package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type GameRouter struct {
	rg *gin.RouterGroup
	db *sqlx.DB
}

func (gr *GameRouter) AddGameRoutes() {
	games := gr.rg.Group("/games")

	games.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "users")
	})
}

func List()
