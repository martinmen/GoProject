package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"gitlab.com/pragmaticreviews/golang-gin-poc/entity"
)

type ExerciseRepository interface {
	Save(exercise entity.Exercise)
	Update(exercise entity.Exercise)
	Delete(exercise entity.Exercise)
	FindAll() []entity.Exercise
	CloseDB()
}

type database struct {
	connection *gorm.DB
}

func NewExerciseRepository() ExerciseRepository {
	db, err := gorm.Open("sqlite3", "test3.db")
	if err != nil {
		panic("Failed to connect database")
	}
	db.AutoMigrate(&entity.Exercise{}, &entity.Exercise{})
	return &database{
		connection: db,
	}
}

func (db *database) CloseDB() {
	err := db.connection.Close()
	if err != nil {
		panic("Fail to close database")
	}
}

func (db *database) Save(exercise entity.Exercise) {
	db.connection.Create(&exercise)
}

func (db *database) Update(exercise entity.Exercise) {
	db.connection.Save(&exercise)
}

func (db *database) Delete(exercise entity.Exercise) {
	db.connection.Delete(&exercise)
}

func (db *database) FindAll() []entity.Exercise {
	var exercises []entity.Exercise
	db.connection.Set("gorm:auto_preload", true).Find(&exercises)
	return exercises
}
