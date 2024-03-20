package cmd

import (
	"github.com/ktr0731/go-fuzzyfinder"
)

func fuzzy(s []string, header string) (string, int, error) {
	if len(s) == 0 {
		return "", -1, nil
	}
	idx, err := fuzzyfinder.Find(
		s,
		func(i int) string {
			return s[i]
		},
		fuzzyfinder.WithHeader(header),
	)
	if err != nil {
		return "", -1, err
	}
	return s[idx], idx, nil
}
