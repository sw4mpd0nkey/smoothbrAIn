package tokenizer

import (
	"fmt"
	"strings"
)

type WordTokenizer struct {
	text         string
	preprocessed []string
}

func NewWordTokenizer(text string) *WordTokenizer {
	return &WordTokenizer{
		text: text,
	}
}

func (word *WordTokenizer) BuildVocabulary() {

	word.preprocessed = strings.Split(word.text, " ")
	for w := range word.preprocessed {
		fmt.Printf("%s\n", word.preprocessed[w])
	}
}
