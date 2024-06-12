package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"server/repository"
	rt "server/router"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	if err := repository.InitDb(); err != nil {
		log.Error().Msg(err.Error())
		return
	}

	router := gin.Default()

	rt.SetupRouter(router)

	srv := &http.Server{
		Addr:    ":80",
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error().Msgf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Info().Msg("Shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Warn().Msgf("Server forced to shutdown: %s", err)
	}

	log.Info().Msg("Server exiting")
}
