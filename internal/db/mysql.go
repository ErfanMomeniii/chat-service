package db

import (
	"database/sql"
	"fmt"
	"github.com/ErfanMomeniii/chat-service/internal/config"

	"github.com/ErfanMomeniii/chat-service/internal/log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Initialize interface {
	CreateDatabase() (*gorm.DB, error)
	Migrate(db *gorm.DB) error
}

type Mysql struct{}

func (m *Mysql) Migrate(db *gorm.DB) error {
	return nil
}

func (m *Mysql) CreateDatabase() (*gorm.DB, error) {
	dbconnect := fmt.Sprintf("%s:%s@tcp(%s:%s)/",
		config.C.Mysql.Username, config.C.Mysql.Password, config.C.Mysql.Host, config.C.Mysql.Port)
	sqldb, err := sql.Open("mysql", dbconnect)

	if err != nil {
		log.Logger.Fatal(err.Error())
	}

	_, err = sqldb.Exec("CREATE DATABASE " + config.C.Mysql.Dbname)

	if err != nil {
		log.Logger.Fatal(err.Error())
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/authform?charset=%s&parseTime=True&loc=Local",
		config.C.Mysql.Username, config.C.Mysql.Password, config.C.Mysql.Host, config.C.Mysql.Port, config.C.Mysql.Charset)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return db, err
	}

	m.Migrate(db)

	return db, nil
}

func Withretry(createDB func() (*gorm.DB, error), limitAttempt int) *gorm.DB {
	for i := 0; i < limitAttempt; i++ {
		db, err := createDB()

		if err != nil {
			log.Logger.Info(err.Error())
		}

		return db
	}

	return nil
}
