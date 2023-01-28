# text-processor
NLP utility repo to interact with text. Initially this is to process English documents, essays, etc


<!-- https://developer.syn.co.in/tutorial/bot/oscova/pretrained-vectors.html -->

# Text Processor

[![Go Report](https://goreportcard.com/badge/github.com/rekram1-node/text-processor)](https://goreportcard.com/report/github.com/rekram1-node/blinkgo) [![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://github.com/rekram1-node/text-processor/blob/main/LICENSE) ![Build Status](https://github.com/rekram1-node/text-processor/actions/workflows/main.yml/badge.svg)

<!-- ![blinkgo](docs/assets/blinkgo-logo.png) -->

Simple library for interacting with blink cameras, mainly: authentication, listing devices/networks/clips, and downloading clips from local storage

This library was made for my purposes but if you would like to see more features open an issue and I will get to it

Credit to MattTW, who's findings: [BlinkMonitorProtocol](https://github.com/MattTW/BlinkMonitorProtocol) I used to create this implementation

## Features

* Extract Sentences and Paragraphs from text
* Show how similar sentences and paragraphs are using word2vec model
* Tokenization, stop word removing, vectorization (done internally)
* Getting key phrases and general sentiment

## Getting Started

### Prerequisites
- [Go](https://go.dev/)
- [Word2vec Model](https://developer.syn.co.in/tutorial/bot/oscova/pretrained-vectors.html)

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
