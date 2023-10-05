package main

import (
	"Sentitube/app"
	"fmt"
)

func main() {
	comments := app.Fetch("")
	fmt.Println(len(comments))
	//app.SentimentAnalysis(comments)

}
