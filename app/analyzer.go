package app

import (
	"github.com/jonreiter/govader"
)

func SentimentAnalysis(c *[]string) (int, int, int, int) {
	var (
		positive int
		neutral  int
		negative int
		not      int
	)
	analyzer := govader.NewSentimentIntensityAnalyzer()

	for _, i := range *c {
		sentiment := analyzer.PolarityScores(i)
		if sentiment.Positive > 0.50 {
			positive++
		} else if sentiment.Neutral > 0.50 {
			neutral++
		} else if sentiment.Negative > 0.50 {
			negative++
		} else {
			not++
		}
	}

	return positive, neutral, negative, not
}
