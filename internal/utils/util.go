package utils

import (
	"github.com/ErfanMomeniii/chat-service/internal/db"
	"github.com/ErfanMomeniii/chat-service/internal/http/request"
	"github.com/ErfanMomeniii/chat-service/internal/model"
	"github.com/ErfanMomeniii/chat-service/internal/repository"
)

func BindToModel(message request.Message) *model.Message {
	userRepo := repository.NewUserRepository(&db.Default{})

	to, _ := userRepo.Get(message.Receiver)
	from, _ := userRepo.Get(message.Sender)

	return &model.Message{
		FromRefer: message.Sender,
		From:      from,
		ToRefer:   message.Receiver,
		To:        to,
		Body:      message.Body,
		IsSeen:    message.IsSeen,
	}
}
