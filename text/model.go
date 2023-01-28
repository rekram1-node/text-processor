package text

import (
	"errors"
	"log"
	"os"

	"code.sajari.com/word2vec"
)

type Word2Vec struct {
	*word2vec.Model
}

func LoadModel(filename string) (*Word2Vec, error) {
	f, err := os.Open(filename)
	if errors.Is(err, os.ErrNotExist) {
		f = getModel(filename)
	}

	defer f.Close()

	w2v, err := word2vec.FromReader(f)

	if err != nil {
		return nil, err
	}

	return &Word2Vec{
		Model: w2v,
	}, nil
}

func getModel(filename string) *os.File {
	f, err := os.Open(filename)

	if err != nil {
		log.Fatalf("failed to open model: %v", err)
	}

	return f
}

func (w2v *Word2Vec) CheckSimilarity(cleanedSourceWordArr, cleanedTargetWordArr []string) (float32, error) {
	sourceExpression := GetExpression(cleanedSourceWordArr)
	targetExpression := GetExpression(cleanedTargetWordArr)

	return w2v.Model.Cos(sourceExpression, targetExpression)
}

func GetExpression(cleanedWordArr []string) word2vec.Expr {
	expr := word2vec.Expr{}
	for _, cleanedWord := range cleanedWordArr {
		expr.Add(1, cleanedWord)
	}
	return expr
}
