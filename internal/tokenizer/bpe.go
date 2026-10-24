package tokenizer

import (
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

func (tokenizer *BPETokenizer) BuildVocabulary() {

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

// Following just stat funcs

func (tokenizer *BPETokenizer) NumberOfWords() int {
	return len(tokenizer.wordFreq)
}

func (tokenizer *BPETokenizer) NumberUniqueChars() int {
	return len(tokenizer.alphabet)
}
