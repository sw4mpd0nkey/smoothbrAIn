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
	maxMerges  int
	corpus     string
	wordFreq   map[string]int
	alphabet   map[rune]int
	pairs      map[Pair]int
	vocabulary map[Pair]string
}

func NewBPETokenizer(text string) *BPETokenizer {

	return &BPETokenizer{
		maxMerges: 1000,
		corpus:    text,
	}
}

func (tokenizer *BPETokenizer) Init() {

	tokenizer.computeWordFreq()
	fmt.Printf("===> We rocking with %d words\n\n", tokenizer.numberOfWords())

	tokenizer.copmuteBaseVocabulary()
	fmt.Printf("===> %d unique characters in total corpus\n", tokenizer.numberUniqueChars())
	fmt.Printf("\n\n")

	tokenizer.computePairFreq()
	fmt.Printf("===> Let's check our Pairs...\n")

	for pair, count := range tokenizer.pairs {
		fmt.Printf("(%s,%s): %d\n", string(pair.A), string(pair.B), count)
	}
	fmt.Printf("\n\n")
}

func (tokenizer *BPETokenizer) TrainBPE() {

	wordSplits := tokenizer.getSplits()

	for step := 0; step < tokenizer.maxMerges; step++ {
		fmt.Println(wordSplits)
	}

}

/*
func (tokenizer *BPETokenizer) mergePair(a, b rune) map[string][]string {
	splits := tokenizer.getSplits()

	for word := range tokenizer.wordFreq {
		split := splits[word]
		if len(split) == 1 {
			continue
		}

		for i := 0; i < len(split)-1; {
			if split[i] == a && split[i+1] == b {
				combo := string(a) + string(b)
				left := split[:i]
				right := split[i+2:]

				split = slices.Concat(left, []string{combo}, right)

			} else {
				i += 1
			}
			splits[word] = split
		}

	}

	return splits
} */

// private funcs

func (tokenizer *BPETokenizer) getSplits() map[string][]rune {
	splits := make(map[string][]rune)

	for word, _ := range tokenizer.wordFreq {
		splits[word] = []rune(word)
	}

	return splits
}

func (tokenizer *BPETokenizer) computeWordFreq() {
	tokenizer.wordFreq = make(map[string]int)
	words := strings.Split(tokenizer.corpus, " ")

	for _, word := range words {
		tokenizer.wordFreq[word]++
	}
}

func (tokenizer *BPETokenizer) computePairFreq() {
	splits := tokenizer.getSplits()
	tokenizer.pairs = make(map[Pair]int)

	for word := range tokenizer.wordFreq {
		split := splits[word]

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

func (tokenizer *BPETokenizer) buildVocabulary() {

	tokenizer.vocabulary = make(map[Pair]string)

	best_pair := Pair{}
	max := 0
	combo := ""

	for pair, freq := range tokenizer.pairs {
		if max < freq {
			best_pair = pair
			max = freq
			combo = string(pair.A) + string(pair.B)
		}
	}

	tokenizer.vocabulary[best_pair] = combo

}

// Following just stat funcs

func (tokenizer *BPETokenizer) numberOfWords() int {
	return len(tokenizer.wordFreq)
}

func (tokenizer *BPETokenizer) numberUniqueChars() int {
	return len(tokenizer.alphabet)
}
