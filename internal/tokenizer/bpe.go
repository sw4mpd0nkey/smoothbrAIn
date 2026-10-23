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
	alphabet := make(map[rune]int)

	for word := range tokenizer.wordFreq {
		for _, char := range word {
			if _, ok := alphabet[char]; !ok {
				alphabet[char] = 1
			}
		}
	}
}

func (tokenizer *BPETokenizer) NumberOfWords() int {
	return len(tokenizer.wordFreq)
}

func (tokenizer *BPETokenizer) NumberUniqueChars() int {
	return len(tokenizer.alphabet)
}
