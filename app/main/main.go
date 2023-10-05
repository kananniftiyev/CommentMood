package main

import (
	"Sentitube/app"
	"fmt"
)

func main() {
	videoID := "bvlFjrcpD6s"
	comments := app.Fetch(videoID)
	fmt.Println(len(comments))
	results := app.SentimentAnalysis(&comments, videoID)
	fmt.Println(results)

}
