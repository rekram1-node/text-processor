package text

import (
	"fmt"
	"strings"

	"github.com/neurosnap/sentences/english"
)

type Text struct {
	Raw        string
	Cleaned    string
	Paragraphs []Paragraph
	// Expressions [][][]word2vec.Expr
}

type Paragraph []Sentence
type Sentence []string

func NewText(rawText string) *Text {
	return &Text{
		Raw: rawText,
	}
}

func (t *Text) Clean(model *Word2Vec) error {
	cleanedText := strings.ReplaceAll(t.Raw, "\n\t", "\n\n")
	paragraphs := strings.Split(cleanedText, "\n\n")

	for _, paragraph := range paragraphs {
		sentences, err := extractSentences(paragraph, model)

		if err != nil {
			return err
		}

		t.Paragraphs = append(t.Paragraphs, sentences)
	}

	return nil
}

func extractSentences(text string, model *Word2Vec) ([]Sentence, error) {
	var extractedSentences []Sentence
	tokenizer, err := english.NewSentenceTokenizer(nil)

	if err != nil {
		return nil, fmt.Errorf("failed to tokenize text %s: %w", text, err)
	}

	for _, sentence := range tokenizer.Tokenize(text) {
		str := sentence.Text
		str = strings.TrimSpace(str)
		cleanedSentence := cleanString(str, model)
		extractedSentences = append(extractedSentences, cleanedSentence)
	}

	return extractedSentences, nil
}

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

func Extract(text string) ([]string, []string, error) {
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
