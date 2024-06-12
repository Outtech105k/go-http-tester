package handler

import (
	"net/http"
	"server/repository"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func POSTPost(c *gin.Context) {
	var json PostRequest
	if err := c.ShouldBindJSON(&json); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, ApiResponce{
			Status: "Internal Server Error",
		})
		log.Error().Msg(err.Error())
		return
	}

	if err := repository.AddPost(json.Body); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, ApiResponce{
			Status: "Internal Server Error",
		})
		log.Error().Msg(err.Error())
		return
	}

	c.IndentedJSON(http.StatusOK, ApiResponce{
		Status: "OK",
	})
}

func GETPosts(c *gin.Context) {
	posts, err := repository.GetPosts()
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, ApiResponce{
			Status: "Internal Server Error",
		})
		return
	}
	c.IndentedJSON(http.StatusOK, ApiResponce{
		Status: "OK",
		Body:   posts,
	})
}

type PostRequest struct {
	Body string `json:"body"`
}
