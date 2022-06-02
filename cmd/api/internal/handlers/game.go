package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ivan-sabo/tic-tac-toe/internal/domain"
	"github.com/ivan-sabo/tic-tac-toe/internal/interfaces"
)

type GameRouter struct {
	rg       *gin.RouterGroup
	gameRepo domain.GameRepository
}

func (gr *GameRouter) AddGameRoutes() {
	games := gr.rg.Group("/games")

	games.GET("/", gr.list)
	games.POST("/", gr.create)
}

func (gr *GameRouter) list(c *gin.Context) {
	games, err := gr.gameRepo.List(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, games)
}

func (gr *GameRouter) create(c *gin.Context) {
	var gameRequest interfaces.NewGameRequest

	c.ShouldBindJSON(&gameRequest)

	board, err := domain.BoardFromString(gameRequest.Board)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	game, err := domain.StartGame(board)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	game, err = gr.gameRepo.Create(c, game)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	r := interfaces.NewGameResponse{
		URL: fmt.Sprintf("%s/games/%s", gr.rg.BasePath(), game.ID),
	}

	c.JSON(http.StatusOK, r)
}
