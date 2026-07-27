package repo

import (
	"fmt"

	"github.com/mateuszmidor/GoStudy/gin-crash-course/entity"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type VideoRepo interface {
	Save(v entity.Video) error
	Update(v entity.Video) error
	Delete(v entity.Video) error
	FindAll() ([]entity.Video, error)
	Close() error
}

type database struct {
	connection *gorm.DB
}

func NewVideoRepo() VideoRepo {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed connecting to test.db: " + err.Error())
	}
	// create database schema
	err = db.AutoMigrate(&entity.Person{}, &entity.Video{})
	if err != nil {
		panic("Failed auto-migration: " + err.Error())
	}
	return &database{connection: db}
}

func (db *database) Close() error {
	fmt.Println("Should close sqlite3 database connection but how?")
	return nil
}

func (db *database) Save(v entity.Video) error {
	return db.connection.Create(&v).Error
}

func (db *database) Update(v entity.Video) error {
	return db.connection.Save(&v).Error
}

func (db *database) Delete(v entity.Video) error {
	return db.connection.Delete(&v).Error
}

func (db *database) FindAll() ([]entity.Video, error) {
	var videos []entity.Video
	err := db.connection.Preload("Author").Find(&videos).Error // fetch videos together with authors
	return videos, err
}
