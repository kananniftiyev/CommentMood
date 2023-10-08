package app

import (
	"os"
)

func GetYoutubeAnalysis(videoID string) (int, int, int, int) {

	if CheckDatabase(videoID) == false {
		comments := Fetch(os.Getenv("API_KEY"), videoID)
		a, b, c, d := SentimentAnalysis(&comments)
		SaveResults(videoID, a, b, c, d)
		return a, b, c, d
	} else {
		return RetrieveResults(videoID)
	}

}

func GetInstagramAnalysis(postID string) {

}
