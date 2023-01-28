# Text Processor

[![Go Report](https://goreportcard.com/badge/github.com/rekram1-node/text-processor)](https://goreportcard.com/report/github.com/rekram1-node/text-processor) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://github.com/rekram1-node/text-processor/blob/main/LICENSE) ![Build Status](https://github.com/rekram1-node/text-processor/actions/workflows/main.yml/badge.svg)


NLP utility library to interact with text documents using a [Word2vec Model](https://developer.syn.co.in/tutorial/bot/oscova/pretrained-vectors.html) Library parses out sentences and paragraphs, removes stop words and tokenizes sentences in order to be consumed by the word2vec comparison functions

## Features

* Extract Sentences and Paragraphs from text
* Show how similar sentences and paragraphs are using word2vec model
* Tokenization, stop word removing, vectorization (done internally)
* Getting key phrases and general sentiment

## Getting Started

### Prerequisites
- [Go](https://go.dev/)
- [Word2vec Model](https://developer.syn.co.in/tutorial/bot/oscova/pretrained-vectors.html) 
- Note: model must be unzipped and in working directory!

## Installing Model

Go to [Word2vec Model](https://developer.syn.co.in/tutorial/bot/oscova/pretrained-vectors.html) and select one of their models to download, I use the 300 Dimension Google News One

### Getting blinkgo

With [Go module](https://github.com/golang/go/wiki/Modules) support, simply add the following import

```go
import "github.com/rekram1-node/text-processor/text"
```

to your code, and then `go [build|run|test]` will automatically fetch the necessary dependencies.

Otherwise, run the following to install the `text-processor` library

```shell
$ go get -u github.com/rekram1-node/text-processor/text
```

## Usage

### Basic Text Comparison

```go
package main

import (
	"fmt"
	"log"

	"github.com/rekram1-node/text-processor/text"
)

func main() {
    t1 := "So much of modern-day life revolves around using opposable thumbs, from holding a hammer to build a home to ordering food delivery on our smartphones. But for our ancestors, the uses were much simpler. Strong and nimble thumbs meant that they could better create and wield tools, stones and bones for killing large animals for food"
    t2 := "A lot of life today involves using opposable thumbs, from using a hammer to build a house to ordering something on our smartphones. But for our predecessors, the uses were much more simple. Powerful and dexterous thumbs meant that they could better make and use tools, stones and bones for killing large animals to eat"

    // load the word2vec model
    m, err := text.LoadModel("yourModelFile")

    if err != nil {
        log.Fatal(err)
    }

    // extract paragraphs and sentences from the text
    t1Paragraphs, t1Sentences, err := text.ExtractAll(t1)
    if err != nil {
        log.Fatal(err)
    }

    // extract paragraphs and sentences from the text
    t2Paragraphs, t2Sentences, err := text.ExtractAll(t2)
    if err != nil {
        log.Fatal(err)
    }

    // compare the two texts and a map of sentences (from 1st document)
    // paired to sentences (from 2nd document) along with a similarity score 
    sim, err := m.MostSimilarSentences(t1Sentences, t2Sentences)
    if err != nil {
        log.Fatal(err)
    }

    // iterate over sentence array and display data
    for _, sentence := range t1Sentences {
        simSentence := sim[sentence]
        if simSentence.Sentence != "" {
            fmt.Println()
            fmt.Println(sentence, "is most similar to:", simSentence.Sentence)
            fmt.Printf("similarity: %v\n", simSentence.Score)
            fmt.Println()
        }
    }

    // compare the two texts and a map of paragraphs (from 1st document)
    // paired to paragraphs (from 2nd document) along with a similarity score 
    simPara, err := m.MostSimilarParagraphs(t1Paragraphs, t2Paragraphs)
    if err != nil {
        log.Fatal(err)
    }

    // iterate over paragraph array and display data
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
```



## Issues

If you have an issue: report it on the [issue tracker](https://github.com/rekram1-node/text-processor/issues)