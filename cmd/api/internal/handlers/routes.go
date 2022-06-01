package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

var (
	router = gin.Default()
)

// getRoutes will create our routes of our entire application
// this way every group of routes can be defined in their own file
// so this one won't be so messy
func GetRouter(db *sqlx.DB) http.Handler {
	v1 := router.Group("/v1")

	gr := GameRouter{
		rg: v1,
		db: db,
	}
	gr.AddGameRoutes()

	return router
}
