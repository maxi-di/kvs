package cmd

import (
	"github.com/ktr0731/go-fuzzyfinder"
)

func fuzzy(s []string, header string) (string, error) {
	idx, err := fuzzyfinder.Find(
		s,
		func(i int) string {
			return s[i]
		},
		fuzzyfinder.WithHeader(header),
	)
	if err != nil {
		return "", err
	}
	return s[idx], nil
}
