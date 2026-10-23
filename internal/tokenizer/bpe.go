package tokenizer

type Pair struct {
	A int
	B int
}

type Merge struct {
	pair   Pair
	second int
}

type BPETokenizer struct {
	corpus string
}

func newBPE(text string) *BPETokenizer {
	return &BPETokenizer{
		corpus: text,
	}
}

func (bpe *BPETokenizer) mergre(ids []int, pair Pair, idx int) []int {
	newIds := make([]int, 0)
	i := 0
	for i < len(ids) {
		if i < len(ids)-1 && ids[i] == pair.A && ids[i+1] == pair.B {
			newIds = append(newIds, idx)
			i += 2
		} else {
			newIds = append(newIds, ids[i])
			i += 1
		}
	}
	return newIds
}

func (bpe *BPETokenizer) getStats(tokens []int) map[Pair]int {
	counts := make(map[Pair]int)
	for i := 0; i < len(tokens)-1; i++ {
		counts[Pair{tokens[i], tokens[i+1]}] += 1
	}
	return counts
}

func (bpe *BPETokenizer) findMostCommonPair(tokens []int) Pair {
	counts := bpe.getStats(tokens)
	mostCommon := Pair{}
	maxCount := -1
	for k, v := range counts {
		if v > maxCount {
			maxCount = v
			mostCommon = k
		}
	}
	return mostCommon
}
