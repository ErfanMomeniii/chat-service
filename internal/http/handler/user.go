package handler

import (
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	GetUser(ctx *gin.Context)
	SaveUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
	UpdateUser(ctx *gin.Context)
}

func GetUser(ctx *gin.Context) {

}

func SaveUser(ctx *gin.Context) {

}

func DeleteUser(ctx *gin.Context) {

}

func UpdateUser(ctx *gin.Context) {

}
