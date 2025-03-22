package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {

	// TODO: Agregar en el

	e := echo.New()

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			_ = next(c)
			log.Info().
				Int("status", c.Response().Status).
				Dur("latency", time.Since(start)).
				Str("client_ip", c.RealIP()).
				Str("method", c.Request().Method).
				Str("path", c.Path()).
				Msg("")
			return nil
		}
	})

	e.GET("/health", func(c echo.Context) error {
		logger := log.With().Str("path", c.Path()).Logger()
		_ = logger.WithContext(c.Request().Context()) // context with sub-logger

		logger.Info().Msg("Health check")
		return c.JSON(http.StatusOK, map[string]string{"status": "OK"})
	})

	e.Logger.Fatal(e.Start(":8080"))
}

func InitLogger(lokiURL string) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnixMs

	logger := log.With().Str("path", "c.Request.URL.Path").Logger()

	logger.Info().Msg("Hello world")
}
