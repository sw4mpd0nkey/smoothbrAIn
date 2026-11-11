package tokenizer

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

type WordTokenizer struct {
	text         string
	preprocessed []string
	vocabulary   map[string]int
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

	sort.Slice(word.preprocessed, func(i, j int) bool {
		return word.preprocessed[i] < word.preprocessed[j]
	})

	word.vocabulary = make(map[string]int)
	vocabCounter := 0
	for _, w := range word.preprocessed {

		if _, exists := word.vocabulary[w]; !exists {
			word.vocabulary[w] = vocabCounter
			vocabCounter++
		}
	}

	fmt.Println("===> Vocabulary build:")
	for key, n := range word.vocabulary {
		fmt.Printf("('%s' : %d)\n", key, n)
	}

}

func trimAllSpace(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
