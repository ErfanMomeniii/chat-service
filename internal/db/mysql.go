package db

import (
	"database/sql"
	"fmt"
	"github.com/ErfanMomeniii/chat-service/internal/config"
	"github.com/ErfanMomeniii/chat-service/internal/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Driver interface {
	Connection
	Register
	Migration
	Rollback
}

type Connection interface {
	Connect(dbname string) (*gorm.DB, error)
}

type Register interface {
	Create(dbname string) error
}

type Migration interface {
	Migrate(db *gorm.DB, models ...interface{}) error
}

type Rollback interface {
	Rollback(db *gorm.DB, models ...interface{}) error
}

type Mysql struct{}

func (m *Mysql) Connect(dbname string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		config.C.Mysql.Username, config.C.Mysql.Password, config.C.Mysql.Host, config.C.Mysql.Port, dbname, config.C.Mysql.Charset)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return db, err
}

func (m *Mysql) Create(dbname string) error {
	conf := fmt.Sprintf("%s:%s@tcp(%s:%s)/",
		config.C.Mysql.Username, config.C.Mysql.Password, config.C.Mysql.Host, config.C.Mysql.Port)

	db, err := sql.Open("mysql", conf)
	if err != nil {
		log.Logger.Fatal(err.Error())
	}

	_, err = db.Exec("CREATE DATABASE " + dbname)
	if err != nil {
		log.Logger.Fatal(err.Error())
	}

	return nil
}

func (m *Mysql) Migrate(db *gorm.DB, models ...interface{}) error {
	err := db.Migrator().AutoMigrate(models)
	if err != nil {
		return err
	}

	return nil
}

func (m *Mysql) Rollback(db *gorm.DB, models ...interface{}) error {
	err := db.Migrator().DropTable(models)
	if err != nil {
		return err
	}

	return nil
}
