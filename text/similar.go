package text

// Structure containing plain sentence and similarity score
type SentAndScore struct {
	// The sentence string
	Sentence string
	// how similar it is to other sentence
	Score float32
}

// Compares two arrays of sentences and returns a hash map of sentence1
// strings paired to sentence2 string with a score similarity as well
func (w2v *Word2Vec) MostSimilarSentences(sentenceArr1, sentenceArr2 []string) (map[string]*SentAndScore, error) {
	mostSim := map[string]*SentAndScore{}
	for _, s1 := range sentenceArr1 {
		mostSim[s1] = &SentAndScore{}
		var biggestSim float32 = 0.00
		cleanedS1 := cleanString(s1, w2v)
		for _, s2 := range sentenceArr2 {
			cleanedS2 := cleanString(s2, w2v)
			sim, err := w2v.CheckSimilarity(cleanedS1, cleanedS2)

			if err != nil {
				return nil, err
			}

			if sim > biggestSim {
				biggestSim = sim
				mostSim[s1] = &SentAndScore{
					Sentence: s2,
					Score:    sim,
				}
			}
		}
	}
	return mostSim, nil
}

// Structure containing plain sentence and similarity score
type ParaAndScore struct {
	// The paragraph string
	Paragraph string
	// how similar it is to other paragraph
	Score float32
}

// Compares two arrays of paragraph and returns a hash map of paragraph1
// strings paired to paragraph2 string with a score similarity as well
func (w2v *Word2Vec) MostSimilarParagraphs(paragraphArr1, paragraphArr2 []string) (map[string]*ParaAndScore, error) {
	mostSim := map[string]*ParaAndScore{}
	for _, p1 := range paragraphArr1 {
		mostSim[p1] = &ParaAndScore{}
		var biggestSim float32 = 0.00
		cleanedP1 := cleanString(p1, w2v)
		for _, p2 := range paragraphArr2 {
			cleanedS2 := cleanString(p2, w2v)
			sim, err := w2v.CheckSimilarity(cleanedP1, cleanedS2)

			if err != nil {
				return nil, err
			}

			if sim > biggestSim {
				biggestSim = sim
				mostSim[p1] = &ParaAndScore{
					Paragraph: p2,
					Score:     sim,
				}
			}
		}
	}

	return mostSim, nil
}
