package router

import (
	"net/http"
	"server/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) error {
	router.POST("/post", handler.POSTPost)
	router.GET("/post", handler.GETPosts)

	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, handler.ApiResponce{
			Status: "Not Found",
		})
	})
	return nil
}
