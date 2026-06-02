package font

import "strings"

func GenerateFont() map[rune][]string {
	mapping := map[rune][]string{}
	for ch := rune(32); ch <= 126; ch++ {

		res := []string{}

		for i := 0; i < 8; i++ {
			res = append(res, strings.Repeat(string(ch), 8))
		}
		mapping[ch] = res
	}
	return mapping

}
