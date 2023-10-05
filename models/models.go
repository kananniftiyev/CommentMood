package models

type Comment struct {
	ID       uint `gorm:"primaryKey;autoIncrement"`
	VideoID  string
	Positive int
	Neutral  int
	Negative int
}
