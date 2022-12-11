package handler

import (
	"github.com/gin-gonic/gin"
)

type MessageHandler interface {
	GetMessage(ctx *gin.Context)
	SendMessage(ctx *gin.Context)
	DeleteMessage(ctx *gin.Context)
	UpdateMessage(ctx *gin.Context)
}

func GetMessage(ctx *gin.Context) {

}

func SendMessage(ctx *gin.Context) {

}

func DeleteMessage(ctx *gin.Context) {

}

func UpdateMessage(ctx *gin.Context) {

}
