package text

import (
	"os"

	"code.sajari.com/word2vec"
)

// Holds loaded Model
type Word2Vec struct {
	*word2vec.Model
}

// loads the word2vec model
func LoadModel(filename string) (*Word2Vec, error) {
	f, err := os.Open(filename)

	if err != nil {
		return nil, err
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

// checks the similarity between two cleaned word arrays
func (w2v *Word2Vec) CheckSimilarity(cleanedSourceWordArr, cleanedTargetWordArr []string) (float32, error) {
	if len(cleanedSourceWordArr) == 0 || len(cleanedTargetWordArr) == 0 {
		return 0.00, nil
	}

	sourceExpression := GetExpression(cleanedSourceWordArr)
	targetExpression := GetExpression(cleanedTargetWordArr)

	return w2v.Model.Cos(sourceExpression, targetExpression)
}

// turns a wordarray into word2vec expression
func GetExpression(cleanedWordArr []string) word2vec.Expr {
	expr := word2vec.Expr{}
	for _, cleanedWord := range cleanedWordArr {
		expr.Add(1, cleanedWord)
	}

	return expr
}
