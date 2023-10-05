package app

import (
	"github.com/jonreiter/govader"
)

func SentimentAnalysis(c *[]string, v string) map[string]int {
	var (
		positive int
		neutral  int
		negative int
		not      int
	)
	results := make(map[string]int)
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
	results["positive"] = positive
	results["neutral"] = neutral
	results["negative"] = negative
	results["not"] = not
	return results
}
