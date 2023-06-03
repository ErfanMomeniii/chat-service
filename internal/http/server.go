package http

import (
	"context"
	"fmt"
	"github.com/ErfanMomeniii/chat-service/internal/app"
	"github.com/ErfanMomeniii/chat-service/internal/config"
	internalHandler "github.com/ErfanMomeniii/chat-service/internal/http/handler"
	"github.com/ErfanMomeniii/chat-service/internal/log"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type server struct {
	engine *gin.Engine
	http   *http.Server
}

func NewServer() *server {
	g := gin.New()
	g.Use(gin.Recovery())
	return &server{
		engine: g,
		http: &http.Server{
			Addr:    config.C.HTTPServer.Listen,
			Handler: g,
		},
	}
}

func (s *server) RegisterRoutes() {
	v1 := s.engine.Group("/v1")
	{
		m := v1.Group("/message")
		{
			messageHandler := internalHandler.NewMessageHandler()
			m.POST("", messageHandler.Send)
			m.GET("/:messageId", messageHandler.Get)
			m.DELETE("/:messageId", messageHandler.Delete)
			m.PUT("/:messageId", messageHandler.Update)
		}

		u := v1.Group("/user")
		{
			userHandler := internalHandler.NewUserHandler()
			u.POST("", userHandler.Save)
			u.GET("/:userId", userHandler.Get)
			u.DELETE("/:userId", userHandler.Delete)
			u.PUT("/:userId", userHandler.Update)
		}

		c := v1.Group("/conversation")
		{
			conversationHandler := internalHandler.NewConversationHandler()
			c.GET("/:fromUserId/:toUserId/messages", conversationHandler.GetMessages)
		}
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
