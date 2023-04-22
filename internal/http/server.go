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
			m.POST("", messageHandler.SendMessage)
			m.GET("/:messageId", messageHandler.GetMessage)
			m.DELETE("/:messageId", messageHandler.DeleteMessage)
			m.PUT("/:messageId", messageHandler.UpdateMessage)
		}

		u := v1.Group("/user")
		{
			userHandler := internalHandler.NewUserHandler()
			u.POST("", userHandler.SaveUser)
			u.GET("/:userId", userHandler.GetUser)
			u.DELETE("/:userId", userHandler.DeleteUser)
			u.PUT("/:userId", userHandler.UpdateUser)
		}

		c := v1.Group("/conversation")
		{
			conversationHandler := internalHandler.NewConversationHandler()
			c.GET("/:fromUserId/:toUserId", conversationHandler.GetMessages)
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
