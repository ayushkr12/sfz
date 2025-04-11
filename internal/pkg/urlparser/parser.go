// Description: This file contains the implementation of the URL parser.
package urlparser

import (
	"fmt"
	"net/url"
	"slices"
)

type FuzzContext struct {
	ParsedURLs     []*URL
	FuzzIdentifier string
}

// NewFuzzContext creates a new FuzzContext with the given fuzz identifier
func NewFuzzContext(fuzzIdentifier string) *FuzzContext {
	return &FuzzContext{FuzzIdentifier: fuzzIdentifier}
}

// UpdateURLsWithFuzzPaths updates the parsed list of URLs with fuzzable paths from the given raw URLs
func (ctx *FuzzContext) UpdateURLsWithFuzzPaths(rawURLs []string) ([]*URL, []error) {
	var errs []error

	for _, rawURL := range rawURLs {
		_, err := ctx.UpdateURLWithFuzzPaths(rawURL)
		if err != nil {
			errs = append(errs, fmt.Errorf("error while updating URL %q: %v", rawURL, err))
		}
	}

	return ctx.ParsedURLs, errs
}

// UpdateURLWithFuzzPaths updates the parsed list of URLs with fuzzable paths from the given raw URL
func (ctx *FuzzContext) UpdateURLWithFuzzPaths(rawURL string) ([]*URL, error) {
	parsedRawURL, err := url.Parse(rawURL)
	if err != nil {
		return nil, fmt.Errorf("error while parsing URL: %v", err)
	}

	baseURL := parsedRawURL.Scheme + "://" + parsedRawURL.Host
	fuzzablePaths := InjectFUZZIdentifiers(parsedRawURL.Path, ctx.FuzzIdentifier)

	// Check if base URL already exists
	for _, p := range ctx.ParsedURLs {
		if baseURL == p.BaseURL {
			for _, fuzzablePath := range fuzzablePaths {
				if !slices.Contains(p.FuzzablePaths, fuzzablePath) {
					p.FuzzablePaths = append(p.FuzzablePaths, fuzzablePath)
				}
			}
			return ctx.ParsedURLs, nil
		}
	}

	// If not found, add new entry and UPDATE ctx.ParsedURLs
	newURL := &URL{
		BaseURL:       baseURL,
		FuzzablePaths: fuzzablePaths,
	}
	ctx.ParsedURLs = append(ctx.ParsedURLs, newURL)

	return ctx.ParsedURLs, nil
}
