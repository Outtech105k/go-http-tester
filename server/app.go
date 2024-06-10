package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	if err := initDb(); err != nil {
		log.Error().Msg(err.Error())
	}

	r := gin.Default()

	r.POST("/post", func(c *gin.Context) {
		var json PostRequest
		if err := c.ShouldBindJSON(&json); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, ApiResponce{
				Status: "Internal Server Error",
			})
			log.Error().Msg(err.Error())
			return
		}

		if err := addPost(json.Body); err != nil {
			c.IndentedJSON(http.StatusInternalServerError, ApiResponce{
				Status: "Internal Server Error",
			})
			log.Error().Msg(err.Error())
			return
		}

		c.IndentedJSON(http.StatusOK, ApiResponce{
			Status: "OK",
		})
	})

	r.GET("/post", func(c *gin.Context) {
		posts, err := getPosts()
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
	})

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, ApiResponce{
			Status: "Not Found",
		})
	})
	r.Run(":80")
}

type PostRequest struct {
	Body string `json:"body"`
}

type ApiResponce struct {
	Status string `json:"status"`
	Body   any    `json:"body"`
}
