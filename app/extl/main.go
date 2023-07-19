package main

import (
	"context"
	"os"
	"os/signal"
	routes "pokemaster-api/interface/api/extl/v1"
	"pokemaster-api/utils/config"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config.LoadConfig()
	e := echo.New()
	e.Use(middleware.CORS())

	// route
	routes.API(e)

	// Start server
	go func() {
		if err := e.Start(":" + os.Getenv("APP_PORT_POKEMASTER")); err != nil {
			e.Logger.Info("Shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
