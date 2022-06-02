package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	games.GET("/:uuid", gr.get)
	games.PUT("/:uuid", gr.put)
}

func (gr *GameRouter) list(c *gin.Context) {
	games, err := gr.gameRepo.List(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, interfaces.NewListGameResponse(games))
}

func (gr *GameRouter) create(c *gin.Context) {
	var gameRequest interfaces.PostGameRequest
	c.ShouldBindJSON(&gameRequest)

	board, err := gameRequest.ToEntity()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	game, err := domain.StartGame(board)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	game, err = gr.gameRepo.Create(c, game)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	r := interfaces.NewGameResponse{
		URL: fmt.Sprintf("%s/games/%s", gr.rg.BasePath(), game.ID),
	}

	c.JSON(http.StatusOK, r)
}

func (gr *GameRouter) get(c *gin.Context) {
	id := c.Param("uuid")

	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, "bad uuid")
		return
	}

	game, err := gr.gameRepo.Get(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, interfaces.NewGetGameReponse(game))
}

func (gr *GameRouter) put(c *gin.Context) {
	id := c.Param("uuid")

	if _, err := uuid.Parse(id); err != nil {
		c.JSON(http.StatusBadRequest, "bad uuid")
		return
	}

	game, err := gr.gameRepo.Get(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	var gameRequest interfaces.PostGameRequest
	c.ShouldBindJSON(&gameRequest)

	board, err := gameRequest.ToEntity()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if err := game.PlayUserMove(board); err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	if err := gr.gameRepo.Update(c, game); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, interfaces.NewPutGameReponse(game))
}
