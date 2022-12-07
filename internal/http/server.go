package http

import (
	"context"
	"github.com/ErfanMomeniii/chat-service/internal/app"
	"github.com/ErfanMomeniii/chat-service/internal/config"
	"github.com/ErfanMomeniii/chat-service/internal/http/handler"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type server struct {
	e *echo.Echo
}

func NewServer() *server {
	e := echo.New()

	e.HideBanner = true
	e.Server.ReadTimeout = config.C.HTTPServer.ReadTimeout
	e.Server.WriteTimeout = config.C.HTTPServer.WriteTimeout
	e.Server.ReadHeaderTimeout = config.C.HTTPServer.ReadHeaderTimeout
	e.Server.IdleTimeout = config.C.HTTPServer.IdleTimeout
	e.Validator = &CustomValidator{Validator: validator.New()}

	return &server{
		e: e,
	}
}

func (s *server) Serve() {
	v1 := s.e.Group("v1")
	v1.POST("/message", handler.SendMessage)
	v1.GET("/message:id", handler.SendMessage)

	go func() {
		if err := s.e.Start(config.C.HTTPServer.Listen); err != nil && err != http.ErrServerClosed {
			s.e.Logger.Fatalf("shutting down the server (%v). err: %v", config.C.HTTPServer.Listen, err)
		}
	}()

	go func() {
		<-app.A.Ctx.Done()
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := s.e.Shutdown(ctx); err != nil {
			s.e.Logger.Fatal(err)
		}
	}()
}
