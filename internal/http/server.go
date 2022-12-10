package http

import (
	"context"
	"fmt"
	"github.com/ErfanMomeniii/chat-service/internal/app"
	"github.com/ErfanMomeniii/chat-service/internal/config"
	"github.com/ErfanMomeniii/chat-service/internal/log"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type server struct {
	http http.Server
}

func NewServer() *server {
	g := gin.New()
	g.Use(gin.Recovery())
	return &server{
		http: http.Server{
			Addr:    config.C.HTTPServer.Listen,
			Handler: g,
		},
	}
}

func (s *server) Serve() {
	go func() {
		if err := s.http.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Logger.Fatal(fmt.Sprintf("shutting down the server (%v). err: %v", config.C.HTTPServer.Listen, err))
		}
	}()

	go func() {
		<-app.A.Ctx.Done()
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := s.http.Shutdown(ctx); err != nil {
			log.Logger.Fatal(err.Error())
		}
	}()
}
