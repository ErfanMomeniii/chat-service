package repository

import (
	"github.com/ErfanMomeniii/chat-service/internal/config"
	"github.com/ErfanMomeniii/chat-service/internal/db"
	"github.com/ErfanMomeniii/chat-service/internal/model"
	"github.com/fatih/structs"
	"gorm.io/gorm"
)

type UserCRUD interface {
	Save(...model.User) error
	Get(userId uint) (model.User, error)
	GetAll() ([]model.User, error)
	Delete(userId uint) error
	Update(user model.User) error
}

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db db.Driver) *UserRepository {
	dbConn, _ := db.Connect(config.C.Mysql.Dbname)

	return &UserRepository{
		DB: dbConn,
	}
}

func (repo *UserRepository) Save(users ...model.User) error {
	var userSlice []model.User
	for _, user := range users {
		userSlice = append(userSlice, user)
	}

	repo.DB.Create(userSlice)

	return nil
}

func (repo *UserRepository) Get(userId uint) (model.User, error) {
	var user model.User
	repo.DB.First(&user, userId)

	return user, nil
}

func (repo *UserRepository) GetAll() ([]model.User, error) {
	var users []model.User
	repo.DB.Find(&users)

	return users, nil
}

func (repo *UserRepository) Delete(userId uint) error {
	repo.DB.Delete(&model.User{}, userId)

	return nil
}

func (repo *UserRepository) Update(user model.User) error {
	repo.DB.Model(&model.User{}).Omit(structs.Name(gorm.Model{})).Updates(&user)

	return nil
}
