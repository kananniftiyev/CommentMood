package models

type Response struct {
	VideoID  string `json:"videoID"`
	Positive int    `json:"positive"`
	Neutral  int    `json:"neutral"`
	Negative int    `json:"negative"`
	Not      int    `json:"not"`
}

type VideoRequest struct {
	VideoID string `json:"videoId"`
}
