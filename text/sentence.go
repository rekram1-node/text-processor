package text

import (
	rake "github.com/afjoseph/RAKE.Go"
	"github.com/jonreiter/govader"
)

// Key Value pair of govader score with phrase
type KeyPhrase struct {
	// govader parsed phrase or word
	Phrase string
	// govader score
	Score float64
}

// Parses out phrases along with their score using govader
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

// Uses RAKE.Go to get the sentiment from text
// responses include: "positive", "neutral", or "negative"
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
