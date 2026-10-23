package tokenizer

import "strings"

type BPETokenizer struct {
	corpus   string
	wordFreq map[string]int
	alphabet map[rune]int
	pairs    map[string]int
}

func NewBPETokenizer(text string) *BPETokenizer {
	return &BPETokenizer{
		corpus: text,
	}
}

func (tokenizer *BPETokenizer) ComputeWordFreq() {
	tokenizer.wordFreq = make(map[string]int)
	words := strings.Split(tokenizer.corpus, " ")

	for _, word := range words {
		tokenizer.wordFreq[word]++
	}
}

func (tokenizer *BPETokenizer) CopmuteBaseVocabulary() {
	tokenizer.alphabet = make(map[rune]int)
	for word := range tokenizer.wordFreq {
		for _, char := range word {
			tokenizer.alphabet[char]++
		}
	}
}

func (tokenizer *BPETokenizer) ComputePairFreq() {

}

// Following just stat funcs

func (tokenizer *BPETokenizer) NumberOfWords() int {
	return len(tokenizer.wordFreq)
}

func (tokenizer *BPETokenizer) NumberUniqueChars() int {
	return len(tokenizer.alphabet)
}
