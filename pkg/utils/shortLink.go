package utils

import (
	"errors"

	"github.com/teris-io/shortid"
)

func GenerateShortLink() (string, error) {
	shortLink, err := shortid.Generate()
	if err != nil {
		return shortLink, errors.New("error generating short link")
	}

	return shortLink, nil
}
