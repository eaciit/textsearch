package textsearch

import (
	"sort"
	"strings"
)

type SimilaritySetting struct {
	Method          string
	Split           bool
	SplitDelimeters []rune
	MinPerWord      int
}

func NewSimilaritySetting() *SimilaritySetting {
	s := new(SimilaritySetting)
	s.Method = "Soundex"
	s.Split = true
	s.SplitDelimeters = []rune{' '}
	s.MinPerWord = 80
	return s
}

func Similarity(s1 string, s2 string, setting *SimilaritySetting) int {
	//-- prepare setting
	if setting == nil {
		setting = NewSimilaritySetting()
	}

	//-- split if required
	var s1s, s2s []string
	if setting.Split == true {
		s1s = Split(s1, setting.SplitDelimeters)
		s2s = Split(s2, setting.SplitDelimeters)
	} else {
		s1s = []string{s1}
		s2s = []string{s2}
	}

	//-- get the sorted checksum
	cs1s := getSimilarityChecksum(s1s)
	cs2s := getSimilarityChecksum(s2s)

	//-- least number of checksum should be first
	if len(cs1s) > len(cs2s) {
		cs1s, cs2s = cs2s, cs1s
	}

	diffs := make([]int, 0)

	//-- iteration
	for _, v1 := range cs1s {
		idxFound := -1

		//-- comparison process
		for i2, v2 := range cs2s {
			diff := EncodedSoundexDiff(v1, v2)
			if diff >= setting.MinPerWord {
				idxFound = i2
				diffs = append(diffs, diff)
				break
			}
		}

		//-- found
		if idxFound >= 0 {
			cs2s = append(cs2s[:idxFound], cs2s[idxFound+1:]...)
		}
	}

	if len(cs2s) > 0 {
		for _ = range cs2s {
			diffs = append(diffs, 0)
		}
	}

	var sums float64
	sums = 0
	for _, v := range diffs {
		sums += float64(v)
	}
	avg := int(sums / float64(len(diffs)))
	return avg
}

func getSimilarityChecksum(s1s []string) []string {
	results := make([]string, 0)
	for _, s := range s1s {
		i := EncodeSoundex(s)
		results = append(results, i)
	}
	sort.Strings(results)
	return results
}

func Split(s string, delims []rune) []string {
	return strings.FieldsFunc(s, func(r rune) bool {
		found := false
		for _, delim := range delims {
			if r == delim {
				return true
			}
		}
		return found
	})
}
