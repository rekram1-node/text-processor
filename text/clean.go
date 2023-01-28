package text

import (
	"regexp"
	"strings"

	"code.sajari.com/word2vec"
	"github.com/bbalet/stopwords"
)

// cleans a string into a word array
func cleanString(s string, model *Word2Vec) []string {
	var words []string
	for _, word := range strings.Split(stopwords.CleanString(s, "EN", false), " ") {
		word = strings.TrimSpace(word)

		if word == "" {
			continue
		}

		word = cleanWord(word)

		if wordInModel(word, model) {
			words = append(words, word)
		}
	}

	return words
}

// trims the word
func cleanWord(s string) string {
	// Remove "St.", for example in "St. Paul"
	newStr := strings.ReplaceAll(s, "St. ", "")
	// ligatures

	newStr = strings.ReplaceAll(newStr, "æ", "ae")

	// Remove hyphens
	newStr = strings.ReplaceAll(newStr, "-", " ")

	// Remove punctuation
	a := regexp.MustCompile(`[,'"“;:”’]`)
	newStr = a.ReplaceAllString(newStr, "")

	// Remove numerics
	numerics := regexp.MustCompile(`\d+`)
	newStr = numerics.ReplaceAllString(newStr, "")

	for _, w := range stopWords {
		re := regexp.MustCompile("(?i)\\b" + w + "\\b")
		newStr = re.ReplaceAllString(newStr, "")
	}

	// strip leading and trailing whitespace
	newStr = strings.TrimSpace(newStr)

	// Compress multiple whitespaces into a single space
	b := regexp.MustCompile(`[\s]{2,}`)
	newStr = b.ReplaceAllString(newStr, " ")

	// lower case
	return strings.ToLower(newStr)
}

// checks if a word is contained in the model
func wordInModel(w string, model *Word2Vec) bool {
	e := word2vec.Expr{}
	e.Add(1, w)
	_, err := e.Eval(model.Model)

	return err == nil
}
