package app

import (
	"Sentitube/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func connectDB() *gorm.DB {
	dsn := "user=postgres password=kanan123 dbname=CommentMood host=localhost port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to Connect to Database")
	}
	return db
}

func init() {
	db := connectDB()
	db.AutoMigrate(&models.Comment{})
	db.Commit()
}

func SaveResults(v string, positive, neutral, negative, not int) {
	db := connectDB()
	comment := models.Comment{VideoID: v, Positive: positive, Neutral: neutral, Negative: negative, Not: not}
	db.Create(&comment)

}

func CheckDatabase(videoID string) bool {
	db := connectDB()
	var exist models.Comment
	if err := db.Where("video_id = ?", videoID).First(&exist).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			log.Fatal("Error checking for existing record:", err)
		}
	} else {
		// Record already exists, you can handle this case as needed
		return true
	}
	return false
}

func RetrieveResults(videoId string) (int, int, int, int) {
	db := connectDB()
	var getComment models.Comment
	db.Where("video_id = ?", videoId).First(&getComment)
	return getComment.Positive, getComment.Neutral, getComment.Negative, getComment.Not
}
