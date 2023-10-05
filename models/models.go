package models

type Comment struct {
	ID       uint `gorm:"primaryKey;autoIncrement"`
	Comment  string
	Positive float64
	Neutral  float64
	Negative float64
}
