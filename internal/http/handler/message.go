package handler

import (
	"github.com/ErfanMomeniii/chat-service/internal/db"
	"github.com/ErfanMomeniii/chat-service/internal/http/request"
	"github.com/ErfanMomeniii/chat-service/internal/http/response"
	"github.com/ErfanMomeniii/chat-service/internal/repository"
	"github.com/ErfanMomeniii/chat-service/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	MessageRepo = repository.NewMessageRepository(&db.Mysql{})
)

type MessageHandler interface {
	GetMessage(ctx *gin.Context)
	SendMessage(ctx *gin.Context)
	DeleteMessage(ctx *gin.Context)
	UpdateMessage(ctx *gin.Context)
}

func GetMessage(ctx *gin.Context) {
	messageId := interface{}(ctx.Param("messageId")).(uint)

	result, err := MessageRepo.Get(messageId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	message := response.Message{
		Receiver: result.To.Username,
		Sender:   result.From.Username,
		Body:     result.Body,
		IsSeen:   result.IsSeen,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": message,
	})
}

func SendMessage(ctx *gin.Context) {
	var message request.Message

	if err := ctx.ShouldBindJSON(&message); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	messageModel := utils.BindToModel(message)

	if err := MessageRepo.Save(*messageModel); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": "success",
	})
}

func DeleteMessage(ctx *gin.Context) {
	messageId := interface{}(ctx.Param("messageId")).(uint)

	err := MessageRepo.Delete(messageId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": "success",
	})
}

func UpdateMessage(ctx *gin.Context) {
	var message request.Message

	messageId := interface{}(ctx.Param("messageId")).(uint)

	if err := ctx.ShouldBindJSON(&message); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	messageModel := utils.BindToModel(message)
	messageModel.ID = messageId

	if err := MessageRepo.Update(*messageModel); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": "success",
	})
}
