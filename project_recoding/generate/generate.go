package generator

func GeneratePattern(c rune) []string {
	pattern := map[rune][]string{
		'A': {
			"  ##  ",
			" #  # ",
			" #  # ",
			" #### ",
			" #  # ",
			" #  # ",
			" #  # ",
			"      ",
		},
		'Z': {
			" #### ",
			"    # ",
			"   #  ",
			"  #   ",
			" #    ",
			" #    ",
			" #### ",
			"      ",
		},
	}
	if c < 'A' || c > 'Z' {
		return []string{}
	}

	if ch, ok := pattern[c]; ok {
		return ch
	}

	return []string{
		"      ",
		"      ",
		"      ",
		"      ",
		"      ",
		"      ",
		"      ",
		"      ",
	}
}
