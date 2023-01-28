package text

import (
	"fmt"

	rake "github.com/afjoseph/RAKE.Go"
	"github.com/jonreiter/govader"
)

func GetKeyWords(s string) {
	candidates := rake.RunRake(s)

	fmt.Println("key words:")
	for _, candidate := range candidates {
		fmt.Println(candidate.Key)
	}
}

func GetSentiment(s string) {
	analyzer := govader.NewSentimentIntensityAnalyzer()
	sentiment := analyzer.PolarityScores(s)

	var sent string

	switch comp := sentiment.Compound; {
	case comp >= 0.05:
		sent = "positive"
	case comp > -0.05 && comp < 0.05:
		sent = "neutral"
	case comp <= -0.05:
		sent = "negative"
	}

	fmt.Println("sentiment:", sent)
}
