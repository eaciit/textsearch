package textsearch

import (
	"fmt"
	"testing"
)

func Test_Soundex(t *testing.T) {
	diff := SoundexDiff("Canon", "Canyon")
	if diff >= 80 {
		fmt.Println("OK, both words are similar")
	} else {
		fmt.Printf("Sorry but both words are different. Similarity index is: %d \n", diff)
	}
}

func Test_TextSimilarity(t *testing.T) {
	s1 := "My name is Arief Darmawan"
	s2 := "Arief Darmawan is the name"
	setting := NewSimilaritySetting()
	setting.SplitDelimeters = []rune{' ', '.', '-'}
	similar := Similarity(s1, s2, setting)

	if diff >= 80 {
		fmt.Println("OK, both sentences are similar")
	} else {
		fmt.Printf("Sorry but both sentences are different. Similarity index is: %d \n", similar)
	}
}
