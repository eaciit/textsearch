package textsearch

mport (
	"strings"
)

func EncodeSoundex(word string) string {
	if word == "" {
		return "0000"
	}
	input := strings.ToLower(word)
	result := strings.ToUpper(input[0:1])
	code := ""
	lastCode := ""
	for _, rune := range input[1:] {
		switch rune {
		case 'b', 'f', 'p', 'v':
			code = "1"
		case 'c', 'g', 'j', 'k', 'q', 's', 'x', 'z':
			code = "2"
		case 'd', 't':
			code = "3"
		case 'l':
			code = "4"
		case 'm', 'n':
			code = "5"
		case 'r':
			code = "6"
		}
		if lastCode != code {
			lastCode = code
			result = result + lastCode
			if len(result) == 4 {
				break
			}
		}
	}
	return result + strings.Repeat("0", 4-len(result))
}

func DifferenceSoundex(word1, word2 string) int {
	sum := differenceSoundex(word1, word2) + differenceSoundex(word2, word1)
	if sum == 0 {
		return 0
	}
	return sum / 2
}

func differenceSoundex(word1, word2 string) int {
	soundex1 := EncodeSoundex(word1)
	soundex2 := EncodeSoundex(word2)
	if soundex1 == soundex2 {
		return 100
	}
	result := 0
	if strings.Index(soundex2, soundex1[1:]) > -1 {
		result = 3
	} else if strings.Index(soundex2, soundex1[2:]) > -1 || strings.Index(soundex2, soundex1[1:3]) > -1 {
		result = 2
	} else {
		if strings.Index(soundex2, soundex1[1:2]) > -1 {
			result = result + 1
		}
		if strings.Index(soundex2, soundex1[2:3]) > -1 {
			result = result + 1
		}
		if strings.Index(soundex2, soundex1[3:4]) > -1 {
			result = result + 1
		}
	}
	if soundex1[0:1] == soundex2[0:1] {
		result = result + 1
	}
	return result * 25
}