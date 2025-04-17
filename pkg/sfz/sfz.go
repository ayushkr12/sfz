package sfz

import (
	"github.com/ayushkr12/sfz/internal/pkg/urlparser"
)

func GenerateFuzzableURLs(rawURLs []string, fuzzIdentifier string) (fuzzableURLS []string, errs []error) {
	ctx := urlparser.NewFuzzContext(fuzzIdentifier)
	parsedURLs, err := ctx.UpdateURLsWithFuzzPaths(rawURLs)
	if err != nil {
		return nil, err
	}
	for _, parsedURL := range parsedURLs {
		for _, fuzzablePath := range parsedURL.FuzzablePaths {
			fuzzableURLS = append(fuzzableURLS, parsedURL.BaseURL+fuzzablePath)
		}
	}
	return fuzzableURLS, nil
}
