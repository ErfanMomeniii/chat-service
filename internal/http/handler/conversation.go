package handler

import (
	"github.com/ErfanMomeniii/chat-service/internal/db"
	"github.com/ErfanMomeniii/chat-service/internal/repository"
	"github.com/gin-gonic/gin"
)

type ConversationHandler interface {
	GetMessages(ctx *gin.Context)
}

type DefaultConversationHandler struct {
	MessageRepository *repository.MessageRepository
}

func NewConversationHandler() *DefaultConversationHandler {
	return &DefaultConversationHandler{
		MessageRepository: repository.NewMessageRepository(&db.Default{}),
	}
}

func (handler *DefaultConversationHandler) GetMessages(ctx *gin.Context) {}
