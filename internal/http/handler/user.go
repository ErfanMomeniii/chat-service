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
	userId := interface{}(ctx.Param("userId")).(uint)

	result, err := handler.UserRepository.Get(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	user := response.User{
		Username:  result.Username,
		Firstname: result.UserInformation.Firstname,
		Lastname:  result.UserInformation.Lastname,
		Tel:       result.UserInformation.Tel,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": user,
	})
}

func (handler *DefaultUserHandler) SaveUser(ctx *gin.Context) {
	var user request.User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userModel := utils.BindUserRequestToModel(user)

	if err := handler.UserRepository.Save(*userModel); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": "success",
	})
}

func (handler *DefaultUserHandler) DeleteUser(ctx *gin.Context) {
	userId := interface{}(ctx.Param("userId")).(uint)

	err := handler.UserRepository.Delete(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": "success",
	})
}

func (handler *DefaultUserHandler) UpdateUser(ctx *gin.Context) {
	var user request.User

	userId := interface{}(ctx.Param("userId")).(uint)

	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userModel := utils.BindUserRequestToModel(user)
	userModel.ID = userId

	if err := handler.UserRepository.Update(*userModel); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": "success",
	})
}
