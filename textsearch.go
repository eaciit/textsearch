package textsearch

import (
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
	s.SplitDelimeters = " "
	s.MinPerWord = 80
	return s
}

func Similarity(s1 string, s2 string, setting *SimilaritySetting) int {
	if setting == nil {
		setting = NewSimilaritySetting()
	}

	var s1s, s2s []string

	if setting.Split == true {
		s1s = Split(s, setting.SplitDelimeters)
		s2s = Split(s, setting.SplitDelimeters)
	} else {
		s1s = []string{s1}
		s2s = []string{s2}
	}
}

func Split(s string, delims []rune) []string {
	return strings.FieldsFunc(s, func(r rune) bool {
		found := false
		i := 0
		for delim := range delims {
			if r == delim {
				return true
			}
		}
		return found
	})
}
