package main

import (
	"errors"
)

func ValidateInput(input string) (rune, error) {
	for _, char := range input {
		if char < 32 || char > 126 {
			return char, errors.New("invalid ascii")
		}
	}
	return 0, nil
}
