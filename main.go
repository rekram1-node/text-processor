//go:build !codeanalysis
// +build !codeanalysis

package main

import (
	"fmt"
	"log"

	"github.com/rekram1-node/text-processor/text"
)

func main() {
	// t := `Prepare your portfolio for what's ahead with expert financial insights from Barron's. Find out where money, markets and business are headed today.`
	var t1 = `So much of modern-day life revolves around using opposable thumbs, from holding a hammer to build a home to ordering food delivery on our smartphones. But for our ancestors, the uses were much simpler. Strong and nimble thumbs meant that they could better create and wield tools, stones and bones for killing large animals for food`
	var t2 = `A lot of life today involves using opposable thumbs, from using a hammer to build a house to ordering something on our smartphones. But for our predecessors, the uses were much more simple. Powerful and dexterous thumbs meant that they could better make and use tools, stones and bones for killing large animals to eat`
	// t2 := ``
	// text.GetKeyWords(t)
	// fmt.Println()
	// text.GetSentiment(t)
	m, err := text.LoadModel("model-bin.bin")

	if err != nil {
		log.Fatal(err)
	}

	t1Paragraphs, t1Sentences, err := text.Extract(t1)
	if err != nil {
		log.Fatal(err)
	}

	t2Paragraphs, t2Sentences, err := text.Extract(t2)
	if err != nil {
		log.Fatal(err)
	}

	sim, err := m.MostSimilarSentences(t1Sentences, t2Sentences)
	if err != nil {
		log.Fatal(err)
	}

	for _, sentence := range t1Sentences {
		simSentence := sim[sentence]
		if simSentence.Sentence != "" {
			fmt.Println()
			fmt.Println(sentence, "is most similar to:", simSentence.Sentence)
			fmt.Printf("similarity: %v\n", simSentence.Score)
			fmt.Println()
		}
	}

	simPara, err := m.MostSimilarParagraphs(t1Paragraphs, t2Paragraphs)
	if err != nil {
		log.Fatal(err)
	}

	for _, para := range t1Paragraphs {
		simParagraph := simPara[para]
		if simParagraph.Paragraph != "" {
			fmt.Println()
			fmt.Println(para, "is most similar to:", simParagraph.Paragraph)
			fmt.Printf("similarity: %v\n", simParagraph.Score)
			fmt.Println()
		}
	}
}
