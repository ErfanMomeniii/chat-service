package repository

import (
	"github.com/ErfanMomeniii/chat-service/internal/config"
	"github.com/ErfanMomeniii/chat-service/internal/db"
	"github.com/ErfanMomeniii/chat-service/internal/model"
	"github.com/fatih/structs"
	"gorm.io/gorm"
)

type MessageCRUD interface {
	Save(...model.Message) error
	Get(messageId uint) (model.Message, error)
	GetAll() ([]model.Message, error)
	GetAllForUserSender(userId uint) ([]model.Message, error)
	GetAllForUserReceiver(userId uint) ([]model.Message, error)
	GetAllForCommunicate(firstUserId uint, secondUserId uint) ([]model.Message, error)
	Delete(messageId uint) error
	Update(message model.Message) error
}

type MessageRepository struct {
	DB *gorm.DB
}

func NewMessageRepo(db db.Driver) *MessageRepository {
	dbConn, _ := db.Connect(config.C.Mysql.Dbname)

	return &MessageRepository{
		DB: dbConn,
	}
}

func (repo *MessageRepository) Create(messages ...model.Message) error {
	var messageSlice []model.Message
	for _, message := range messages {
		messageSlice = append(messageSlice, message)
	}

	repo.DB.Create(messageSlice)

	return nil
}

func (repo *MessageRepository) Get(messageId uint) (model.Message, error) {
	var message model.Message
	repo.DB.First(&message, messageId)

	return message, nil
}

func (repo *MessageRepository) GetAll() ([]model.Message, error) {
	var messages []model.Message
	repo.DB.Find(&messages)

	return messages, nil
}

func (repo *MessageRepository) GetAllForUserSender(userId uint) ([]model.Message, error) {
	var messages []model.Message
	repo.DB.Where("FromRefer = ?", userId).Find(&messages)

	return messages, nil
}

func (repo *MessageRepository) GetAllForUserReceiver(userId uint) ([]model.Message, error) {
	var messages []model.Message
	repo.DB.Where("ToRefer = ?", userId).Find(&messages)

	return messages, nil
}

func (repo *MessageRepository) GetAllForCommunicate(firstUserId uint, secondUserId uint) ([]model.Message, error) {
	var messages []model.Message
	repo.DB.Where("ToRefer = ? AND FromRefer = ?", firstUserId, secondUserId).Or("ToRefer = ? AND FromRefer = ?", secondUserId, firstUserId).Find(&messages)

	return messages, nil
}

func (repo *MessageRepository) Delete(messageId uint) error {
	repo.DB.Delete(&model.Message{}, messageId)

	return nil
}

func (repo *MessageRepository) Update(message model.Message) error {
	repo.DB.Model(&model.Message{}).Omit(structs.Name(gorm.Model{})).Updates(&message)

	return nil
}
