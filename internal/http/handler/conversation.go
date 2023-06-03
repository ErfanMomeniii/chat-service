package handler

import (
	"github.com/ErfanMomeniii/chat-service/internal/db"
	"github.com/ErfanMomeniii/chat-service/internal/http/response"
	"github.com/ErfanMomeniii/chat-service/internal/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ConversationHandler interface {
	GetMessages(ctx *gin.Context)
}

type DefaultConversationHandler struct {
	MessageRepository *repository.MessageRepository
}

func NewConversationHandler() *DefaultConversationHandler {
	return &DefaultConversationHandler{
		MessageRepository: repository.NewMessageRepository(&db.Mysql{}),
	}
}

func (handler *DefaultConversationHandler) GetMessages(ctx *gin.Context) {
	fromUserId := interface{}(ctx.Param("fromUserId")).(uint)
	toUserId := interface{}(ctx.Param("toUserId")).(uint)

	result, err := handler.MessageRepository.GetAllForCommunicate(fromUserId, toUserId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	var messages []response.Message
	for _, message := range result {
		messages = append(messages, response.Message{
			Receiver: message.To.Username,
			Sender:   message.From.Username,
			Body:     message.Body,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": messages,
	})
}

func (handler *DefaultConversationHandler) GetAll(ctx *gin.Context) {
	fromUserId := interface{}(ctx.Param("fromUserId")).(uint)
	toUserId := interface{}(ctx.Param("toUserId")).(uint)

	result, err := handler.MessageRepository.GetAllForCommunicate(fromUserId, toUserId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	var messages []response.Message
	for _, message := range result {
		messages = append(messages, response.Message{
			Receiver: message.To.Username,
			Sender:   message.From.Username,
			Body:     message.Body,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": messages,
	})
}
