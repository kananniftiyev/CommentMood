package app

import (
	"Sentitube/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func init() {
	dsn := "user=postgres password=kanan123 dbname=CommentMood host=db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to Connect to Database")
	}
	db.AutoMigrate(&models.Comment{})
	db.Commit()
}

func SaveResults(v string, positive, neutral, negative, not int) {
	dsn := "user=postgres password=kanan123 dbname=CommentMood host=db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Could not connect to Database")
	}
	// Check if a record with the same VideoID exist
	var existingComment models.Comment
	if err := db.Where("video_id = ?", v).First(&existingComment).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			log.Fatal("Error checking for existing record:", err)
		}
	} else {
		// Record already exists, you can handle this case as needed
		log.Printf("Record with VideoID %s already exists", v)
		return
	}
	comment := models.Comment{VideoID: v, Positive: positive, Neutral: neutral, Negative: negative, Not: not}
	db.Create(&comment)

}

func CheckDatabase(videoID string) bool {
	dsn := "user=postgres password=kanan123 dbname=CommentMood host=db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Could not connect to Database")
	}
	var exist models.Comment
	if err := db.Where("video_id = ?", videoID).First(&exist).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			log.Fatal("Error checking for existing record:", err)
		}
	} else {
		// Record already exists, you can handle this case as needed
		log.Printf("Record with VideoID %s already exists", videoID)
		return true
	}
	return false
}

func RetrieveResults(videoId string) (int, int, int, int) {
	dsn := "user=postgres password=kanan123 dbname=CommentMood host=db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Could not connect to Database")
	}
	var getComment models.Comment
	db.Where("video_id = ?", videoId).First(&getComment)
	return getComment.Positive, getComment.Neutral, getComment.Negative, getComment.Not
}
