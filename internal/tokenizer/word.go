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
	specialChar := regexp.MustCompile(`([.,?!])`)
	inFlight := splitter.Split(word.text, -1)
	for _, w := range inFlight {
		if specialChar.Match([]byte(w)) {
			word.preprocessed = append(word.preprocessed, specialChar.Split(w, -1)[0])
			word.preprocessed = append(word.preprocessed, specialChar.FindString(w))
		} else {
			word.preprocessed = append(word.preprocessed, w)
		}
	}

	for _, w := range word.preprocessed {
		fmt.Printf("%s\n", w)
	}

}
