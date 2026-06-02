package converter

import "strings"

func StringToArt(input string) string {
	pattern := map[rune][]string{
		'0': {
			" ___ ",
			"|   |",
			"|   |",
			"|   |",
			"|___|",
		},
		'1': {
			"  |  ",
			"  |  ",
			"  |  ",
			"  |  ",
			"  |  ",
		},
		'2': {
			" ___ ",
			"    |",
			" ___|",
			"|    ",
			"|___ ",
		},
	}
	if input == "" {
		return ""
	}

	res := ""

	for _, line := range strings.Split(input, "\n") {
		if line == "" {
			return ""
		}

		pat := make([][]string, 0, len(line))

		for _, ch := range line {
			char, ok := pattern[ch]
			if !ok {
				return ""
			}
			pat = append(pat, char)
		}

		for row := 0; row < 5; row++ {
			for _, char := range pat {
				res += char[row]
			}
			res += "\n"
		}
	}
	return res
}
