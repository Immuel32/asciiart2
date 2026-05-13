package main

import (
	"strings"
)

func GenerateArt(input string, banner map[rune][]string) string {
	if input == "" {
		return ""
	}
	lines := strings.Split(input, "\n")
	result := []string{}

	hasArt := false

	for _, line := range lines {
		if line == "" {
			result = append(result, "")
			continue
		}
		hasArt = true

		asciiLines := make([]string, 8)

		for _, char := range line {
			art, ok := banner[char]
			if !ok {
				art = banner[' ']
			}
			for i := 0; i < 8; i++ {
				asciiLines[i] += art[i]

			}
		}
		result = append(result, asciiLines...)
	}
	output := strings.Join(result, "\n")
	if hasArt {
		output += "\n"
	}
	return output
}
