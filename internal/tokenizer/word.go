package tokenizer

import (
	"fmt"
	"regexp"
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
	splitter := regexp.MustCompile(`(\s)`)
	word.preprocessed = splitter.Split(word.text, -1)
	for w := range word.preprocessed {
		fmt.Printf("%s\n", word.preprocessed[w])

	}
}
