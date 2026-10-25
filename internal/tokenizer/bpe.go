package tokenizer

import (
	"fmt"
	"strings"
)

type MAX_MERGES int

type Pair struct {
	A rune
	B rune
}

type BPETokenizer struct {
	vocabSize  int
	corpus     string
	wordFreq   map[string]int
	alphabet   map[rune]int
	pairs      map[Pair]int
	vocabulary map[Pair]string
	splits     map[string][]rune
}

func NewBPETokenizer(text string) *BPETokenizer {

	return &BPETokenizer{
		vocabSize: 50,
		corpus:    text,
	}
}

func (tokenizer *BPETokenizer) Init() {

	tokenizer.computeWordFreq()
	fmt.Printf("===> We rocking with %d words\n\n", tokenizer.numberOfWords())

	tokenizer.copmuteBaseVocabulary()
	fmt.Printf("===> %d unique characters in total corpus\n", tokenizer.numberUniqueChars())
	fmt.Printf("\n\n")

	tokenizer.generateSplits()
	fmt.Printf("===> Generated word splits\n\n")

	tokenizer.computePairFreq()
	fmt.Printf("===> Let's check our Pairs...\n")

	for pair, count := range tokenizer.pairs {
		fmt.Printf("(%s,%s): %d\n", string(pair.A), string(pair.B), count)
	}
	fmt.Printf("\n\n")

}

func (tokenizer *BPETokenizer) TrainBPE() {

	for step := 0; step < tokenizer.vocabSize; step++ {
		tokenizer.computePairFreq()
		bestPair := tokenizer.findBestPair()
		tokenizer.mergePair(bestPair)
		tokenizer.vocabulary[bestPair] = string(bestPair.A) + string(bestPair.B)
	}

}

// private funcs

func (tokenizer *BPETokenizer) generateSplits() {
	tokenizer.splits = make(map[string][]rune)

	for word := range tokenizer.wordFreq {
		tokenizer.splits[word] = []rune(word)
	}
}

func (tokenizer *BPETokenizer) computeWordFreq() {
	tokenizer.wordFreq = make(map[string]int)
	words := strings.Split(tokenizer.corpus, " ")

	for _, word := range words {
		tokenizer.wordFreq[word]++
	}
}

func (tokenizer *BPETokenizer) computePairFreq() {
	tokenizer.pairs = make(map[Pair]int)

	for word := range tokenizer.wordFreq {
		split := tokenizer.splits[word]

		if len(split) == 1 {
			continue
		}

		for i := range len(split) - 1 {
			pair := Pair{
				A: split[i],
				B: split[i+1],
			}
			tokenizer.pairs[pair]++
		}

	}
}

func (tokenizer *BPETokenizer) copmuteBaseVocabulary() {
	tokenizer.alphabet = make(map[rune]int)
	for word := range tokenizer.wordFreq {
		for _, char := range word {
			tokenizer.alphabet[char]++
		}
	}
}

func (tokenizer *BPETokenizer) findBestPair() Pair {

	best_pair := Pair{}
	max := 0

	for pair, freq := range tokenizer.pairs {
		if max < freq {
			best_pair = pair
			max = freq
		}
	}

	return best_pair

}

func (tokenizer *BPETokenizer) mergePair(pair Pair) {

	for word := range tokenizer.wordFreq {

		split := tokenizer.splits[word]
		if len(split) == 1 {
			continue
		}

		for i := 1; i < len(split)-1; {
			if split[i] == pair.A && split[i+1] == pair.B {

			}
		}
	}
}

// Following just stat funcs

func (tokenizer *BPETokenizer) numberOfWords() int {
	return len(tokenizer.wordFreq)
}

func (tokenizer *BPETokenizer) numberUniqueChars() int {
	return len(tokenizer.alphabet)
}
