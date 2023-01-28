package text

import (
	"fmt"
	"strings"

	"github.com/neurosnap/sentences/english"
)

// Extract all sentences from raw document
func ExtractSentences(text string) ([]string, error) {
	var extractedSentences []string
	tokenizer, err := english.NewSentenceTokenizer(nil)

	if err != nil {
		return nil, fmt.Errorf("failed to tokenize text %s: %w", text, err)
	}

	for _, sentence := range tokenizer.Tokenize(text) {
		str := sentence.Text
		str = strings.TrimSpace(str)
		extractedSentences = append(extractedSentences, str)
	}

	return extractedSentences, nil
}

// Extract all paragraphs and sentences from raw document
func ExtractAll(text string) ([]string, []string, error) {
	var sentenceArr []string
	cleanedText := strings.ReplaceAll(text, "\n\t", "\n\n")
	paragraphArr := strings.Split(cleanedText, "\n\n")

	for _, paragraph := range paragraphArr {
		sentences, err := ExtractSentences(paragraph)
		if err != nil {
			return nil, nil, err
		}
		sentenceArr = append(sentenceArr, sentences...)
	}

	return paragraphArr, sentenceArr, nil
}
