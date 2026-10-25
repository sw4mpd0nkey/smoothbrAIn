package tokenizer

import (
	"fmt"
	"slices"
	"strings"
)

type Pair struct {
	A string
	B string
}

type BPETokenizer struct {
	corpus     string
	wordFreq   map[string]int
	alphabet   map[rune]int
	Pairs      map[Pair]int
	Vocabulary map[Pair]string
}

func NewBPETokenizer(text string) *BPETokenizer {
	return &BPETokenizer{
		corpus: text,
	}
}

func (tokenizer *BPETokenizer) Init() {

	tokenizer.computeWordFreq()
	fmt.Printf("===> We rocking with %d words\n", tokenizer.numberOfWords())
	fmt.Printf("\n\n")

	tokenizer.copmuteBaseVocabulary()
	fmt.Printf("===> %d unique characters in total corpus\n", tokenizer.numberUniqueChars())
	fmt.Printf("\n\n")

	tokenizer.computePairFreq()
	fmt.Printf("===> Let's check our Pairs...\n")

	for pair, count := range tokenizer.Pairs {
		fmt.Printf("(%c,%c): %d\n", pair.A, pair.B, count)
	}
	fmt.Printf("\n\n")
}

func (tokenizer *BPETokenizer) TrainBPE() {
	tokenizer.buildVocabulary()
	fmt.Printf("===> Now lets pump our vocab fam:\n")
	for pair, token := range tokenizer.Vocabulary {
		fmt.Printf("(%c,%c): %s", pair.A, pair.B, token)
	}
	fmt.Printf("\n\n")
}

func (tokenizer *BPETokenizer) mergePair(a, b string) map[string][]string {
	splits := tokenizer.getSplits()

	for word, _ := range tokenizer.wordFreq {
		split := splits[word]
		if len(split) == 1 {
			continue
		}

		for i := 0; i < len(split)-1; {
			if split[i] == a && split[i+1] == b {
				combo := a + b
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
}

func (tokenizer *BPETokenizer) Tokenize() {

}

// private funcs

func (tokenizer *BPETokenizer) getSplits() map[string][]string {
	splits := make(map[string][]string)

	for word, _ := range tokenizer.wordFreq {
		rune_array := []rune(word)
		string_array := make([]string, len(rune_array))
		for letter := range rune_array {
			string_array = append(string_array, string(letter))
		}
		splits[word] = string_array
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
	tokenizer.Pairs = make(map[Pair]int)
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
			tokenizer.Pairs[pair]++
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

	tokenizer.Vocabulary = make(map[Pair]string)

	best_pair := Pair{}
	max := 0
	combo := ""

	for pair, freq := range tokenizer.Pairs {
		if max < freq {
			best_pair = pair
			max = freq
			combo = string(pair.A) + string(pair.B)
		}
	}

	tokenizer.Vocabulary[best_pair] = combo

}

// Following just stat funcs

func (tokenizer *BPETokenizer) numberOfWords() int {
	return len(tokenizer.wordFreq)
}

func (tokenizer *BPETokenizer) numberUniqueChars() int {
	return len(tokenizer.alphabet)
}
