package main

import (
	"errors"
)

func ValidateBanner(banner map[rune][]string) error {
	if banner == nil {
		return errors.New("banner is nil")
	}
	if len(banner) == 0 {
		return errors.New("empty bannerfile")
	}
	if len(banner) != 95 {
		return errors.New("invalid file")
	}
	for r, art := range banner {
		if len(art) != 8 {
			return errors.New("'A' got: invalid content of art")
		}
		if r < 32 || r > 126 {
			return errors.New("invalid file 'A'")
		}
	}
	return nil
}
