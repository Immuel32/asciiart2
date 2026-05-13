package main

import (
	"strings"
)

func SplitInput(input string) []string {
	split := strings.Split(input, "\\n")
	return split
}
