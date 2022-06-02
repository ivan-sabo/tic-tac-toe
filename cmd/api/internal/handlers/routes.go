package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ivan-sabo/tic-tac-toe/internal/domain"
)

var (
	router = gin.Default()
)

// getRoutes will create our routes of our entire application
// this way every group of routes can be defined in their own file
// so this one won't be so messy
func GetRouter(gameRepo domain.GameRepository) http.Handler {
	v1 := router.Group("/v1")

	gr := GameRouter{
		rg:       v1,
		gameRepo: gameRepo,
	}
	gr.AddGameRoutes()

	return router
}
