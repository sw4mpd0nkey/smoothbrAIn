package tokenizer

import "strings"

type BPETokenizer struct {
	corpus   string
	wordFreq map[string]int
	alphabet map[rune]int
}

func NewBPETokenizer(text string) *BPETokenizer {
	return &BPETokenizer{
		corpus:   text,
		wordFreq: make(map[string]int),
		alphabet: make(map[rune]int),
	}
}

func (tokenizer *BPETokenizer) ComputeWordReq() {
	words := strings.Split(tokenizer.corpus, " ")

	for _, word := range words {
		tokenizer.wordFreq[word]++
	}
}

func (tokenizer *BPETokenizer) CopmuteBaseVocabulary() {
	for word := range tokenizer.wordFreq {
		for _, char := range word {
			tokenizer.alphabet[char]++
		}
	}
}

// Following just stat funcs

func (tokenizer *BPETokenizer) NumberOfWords() int {
	return len(tokenizer.wordFreq)
}

func (tokenizer *BPETokenizer) NumberUniqueChars() int {
	return len(tokenizer.alphabet)
}
