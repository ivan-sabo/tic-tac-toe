package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ivan-sabo/tic-tac-toe/internal/domain"
)

type GameRouter struct {
	rg       *gin.RouterGroup
	gameRepo domain.GameRepository
}

func (gr *GameRouter) AddGameRoutes() {
	games := gr.rg.Group("/games")

	games.GET("/", gr.list)
}

func (gr *GameRouter) list(c *gin.Context) {
	games, err := gr.gameRepo.List(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, games)
}
