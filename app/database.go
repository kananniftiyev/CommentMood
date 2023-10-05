package app

import (
	"Sentitube/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func init() {
	dsn := "user=postgres password=kanan123 dbname=CommentMood host=localhost port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to Connect to Database")
	}
	db.AutoMigrate(&models.Comment{})
	db.Commit()
}

func SaveResults(c string, positive, neutral, negative float64) {
	dsn := "user=postgres password=kanan123 dbname=CommentMood host=localhost port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Could not connect to Database")
	}
	comment := models.Comment{Comment: c, Positive: positive, Neutral: neutral, Negative: negative}
	db.Create(&comment)

}
