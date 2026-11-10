package tokenizer

import (
	"fmt"
	"regexp"
	"strings"
)

type WordTokenizer struct {
	text         string
	preprocessed []string
	vocabulary   []string
}

func NewWordTokenizer(text string) *WordTokenizer {
	return &WordTokenizer{
		text: text,
	}
}

func (word *WordTokenizer) BuildVocabulary() {
	splitter := regexp.MustCompile(`(\s)`)
	specialChar := regexp.MustCompile(`([.,?!"':;\-_])`)

	// initial split on white space
	inFlight := splitter.Split(word.text, -1)
	// additional processing for special characters, also removes any additional whitespace
	for _, w := range inFlight {
		if specialChar.Match([]byte(w)) {
			word.preprocessed = append(word.preprocessed, trimAllSpace(specialChar.Split(w, -1)[0]))
			word.preprocessed = append(word.preprocessed, trimAllSpace(specialChar.FindString(w)))
		} else {
			word.preprocessed = append(word.preprocessed, trimAllSpace(w))
		}
	}

	// print out preprocessed
	fmt.Println("===> All words and special characters:")
	for _, w := range word.preprocessed {
		fmt.Printf("%s\n", w)
	}

}

func trimAllSpace(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
