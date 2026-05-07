package main

import "strings"

func RenderLine(text string, banner map[rune][]string) []string {
	res := make([]string, 8)
	for row := 0; row < 8; row++ {
		var str strings.Builder
		for _, char := range text {
			str.WriteString(banner[char][row])
		}
		res[row] = str.String()
	}
	return res
}
