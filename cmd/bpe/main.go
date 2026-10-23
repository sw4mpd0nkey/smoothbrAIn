package main

import (
	"fmt"

	"github.com/sw4mpd0nkey/smoothbrain/internal/tokenizer"
)

func main() {
	text := " Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi lobortis suscipit sapien, eget elementum urna accumsan quis. Sed sed neque non turpis molestie aliquam. Curabitur euismod congue risus, sed tristique nibh venenatis id. Quisque blandit commodo finibus. Sed tincidunt ante ac sem consequat, ac porta ipsum tempor. Vestibulum libero dui, auctor a sollicitudin a, consectetur sit amet sem. Etiam sed tortor massa. Nunc malesuada ligula vitae ex tincidunt, ac sollicitudin nibh pharetra. Duis id arcu volutpat, posuere tellus sed, maximus odio. In lorem elit, posuere interdum venenatis nec, egestas a dui. Etiam sodales risus in convallis dapibus. Nulla id venenatis odio. Cras lobortis nisl ex. Donec tincidunt consectetur urna. Phasellus vitae nisl sed odio lobortis scelerisque. Donec non luctus dolor. Maecenas lacinia quis est eget imperdiet. Mauris sodales, metus quis scelerisque elementum, metus odio molestie orci, vel ornare erat neque eu augue. Maecenas sapien erat, malesuada ut accumsan ac, pharetra cursus lorem. Sed malesuada purus ac libero dictum, ut sodales turpis malesuada. Ut mauris nisi, viverra non rhoncus et, dictum et nisi. Fusce in odio mi. Cras finibus enim at mi aliquet placerat. Donec id turpis neque. Pellentesque fringilla sollicitudin enim, id maximus ligula interdum eget. Duis arcu quam, fermentum in sapien nec, ullamcorper auctor justo. Vestibulum rutrum augue non diam accumsan, eu tincidunt leo vulputate. "
	fmt.Printf("===> Chopping up the following:\n%s\n", text)

	tokenizer := tokenizer.NewBPETokenizer(text)

	tokenizer.ComputeWordFreq()
	fmt.Printf("===> We rocking with %d words\n", tokenizer.NumberOfWords())
	fmt.Printf("\n\n")

	tokenizer.CopmuteBaseVocabulary()
	fmt.Printf("===> %d unique characters in total corpus\n", tokenizer.NumberUniqueChars())
	fmt.Printf("\n\n")

	tokenizer.ComputePairFreq()
	fmt.Printf("===> Let's check our Pairs...\n")

	for pair, count := range tokenizer.Pairs {
		fmt.Printf("(%c,%c): %d\n", pair.A, pair.B, count)
	}
	fmt.Printf("\n\n")

	tokenizer.BuildVocabulary()
	fmt.Printf("===> Now lets pump our vocab fam:\n")
	for pair, token := range tokenizer.Vocabulary {
		fmt.Printf("(%c,%c): %s", pair.A, pair.B, token)
	}
	fmt.Printf("\n\n")
}
