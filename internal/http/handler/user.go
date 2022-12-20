package handler

import (
	"github.com/ErfanMomeniii/chat-service/internal/db"
	"github.com/ErfanMomeniii/chat-service/internal/repository"
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	GetUser(ctx *gin.Context)
	SaveUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
}

type DefaultUserHandler struct {
	UserRepository *repository.UserRepository
}

func NewUserHandler() UserHandler {
	return &DefaultUserHandler{
		UserRepository: repository.NewUserRepository(&db.Default{}),
	}
}

func (handler *DefaultUserHandler) GetUser(ctx *gin.Context) {

}

func (handler *DefaultUserHandler) SaveUser(ctx *gin.Context) {

}

func (handler *DefaultUserHandler) DeleteUser(ctx *gin.Context) {

}

func (handler *DefaultUserHandler) UpdateUser(ctx *gin.Context) {

}
