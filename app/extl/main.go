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
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))

	// route
	routes.API(e)

	// Start server
	go func() {
		port := os.Getenv("PORT")
		if port == "" {
			port = "5000" // Default to port 5000 if PORT is not set
		}

		if err := e.Start(":" + port); err != nil {
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
