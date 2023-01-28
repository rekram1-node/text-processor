package text

import (
	rake "github.com/afjoseph/RAKE.Go"
	"github.com/jonreiter/govader"
)

type KeyPhrase struct {
	Phrase string
	Score  float64
}

func GetKeyPhrases(s string) []KeyPhrase {
	phrases := []KeyPhrase{}
	candidates := rake.RunRake(s)
	for _, candidate := range candidates {
		phrases = append(phrases, KeyPhrase{
			Phrase: candidate.Key,
			Score:  candidate.Value,
		})
	}

	return phrases
}

func GetSentiment(s string) string {
	analyzer := govader.NewSentimentIntensityAnalyzer()
	sentiment := analyzer.PolarityScores(s)

	switch comp := sentiment.Compound; {
	case comp >= 0.05:
		return "positive"
	case comp > -0.05 && comp < 0.05:
		return "neutral"
	case comp <= -0.05:
		return "negative"
	default:
		return "neutral"
	}
}
