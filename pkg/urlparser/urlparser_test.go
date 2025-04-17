package urlparser

import (
	"testing"
)

var rawURLS = []string{"https://example.com/path/to/resource", "https://test.io/another/path"}

func TestGenerateFuzzableURLs(t *testing.T) {
	fuzzIdentifier := "FUZZ"
	fuzzableUrls, errs := GenerateFuzzableURLs(rawURLS, fuzzIdentifier)

	for _, u := range fuzzableUrls {
		t.Log(u)
	}

	if errs != nil {
		t.Logf("Error occured while generating fuzzable URLs: %v", errs)
	}

}
