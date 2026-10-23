package tokenizer

import "strings"

type Pair struct {
	A rune
	B rune
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
	for word, _ := range tokenizer.wordFreq {
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
			combo = string(pair.A + pair.B)
		}
	}

	tokenizer.Vocabulary[best_pair] = combo

}

// private funcs

func (tokenizer *BPETokenizer) getSplits() map[string][]rune {
	splits := make(map[string][]rune)

	for word, _ := range tokenizer.wordFreq {
		splits[word] = []rune(word)
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
