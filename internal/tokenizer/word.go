package tokenizer

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

var splitter = regexp.MustCompile(`(\s)`)
var specialChar = regexp.MustCompile(`([.,?!"':;\-_])`)

var unkownToken = "<|unk|>"
var eot = "<|endoftext|>"

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

	word.vocabulary[unkownToken] = vocabCounter
	vocabCounter++
	word.vocabulary[eot] = vocabCounter

	fmt.Println("===> Vocabulary build:")
	for key, n := range word.vocabulary {
		fmt.Printf("('%s' : %d)\n", key, n)
	}

}

func (word *WordTokenizer) Encode(text string) []int {
	toReturn := []int{}
	wordArray := []string{}

	// initial split on white space
	inFlight := splitter.Split(text, -1)
	// additional processing for special characters, also removes any additional whitespace
	for _, w := range inFlight {
		if specialChar.Match([]byte(w)) {
			wordArray = append(wordArray, trimAllSpace(specialChar.Split(w, -1)[0]))
			wordArray = append(wordArray, trimAllSpace(specialChar.FindString(w)))
		} else {
			wordArray = append(wordArray, trimAllSpace(w))
		}
	}

	for _, w := range wordArray {

		if _, exists := word.vocabulary[w]; exists {
			toReturn = append(toReturn, word.vocabulary[w])
		} else {
			toReturn = append(toReturn, word.vocabulary[unkownToken])
		}
	}

	return toReturn
}

func (word *WordTokenizer) Decode(ids []int) string {
	toReturn := ""

	for i := range ids {
		word := mapLookUp(ids[i], word.vocabulary)
		toReturn = toReturn + " " + word
	}

	return toReturn
}

func trimAllSpace(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func mapLookUp(myValue int, m map[string]int) string {
	for key, value := range m {
		if value == myValue {
			return key
		}
	}

	return unkownToken
}
