package utils

import "unicode"

func IsAvailable(alpha []string, str string) bool {
	for i := 0; i < len(alpha); i++ {
		if alpha[i] == str {
			return true
		}
	}
	return false
}

func isVietnameseLetter(char rune) bool {
	vietnameseCharset := []*unicode.RangeTable{
		unicode.Letter,
	}

	for _, charset := range vietnameseCharset {
		if unicode.In(char, charset) {
			return true
		}
	}

	return false
}

