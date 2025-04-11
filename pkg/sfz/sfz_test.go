package sfz

import (
	"testing"
)

func TestGenerateFuzzableURLs(t *testing.T) {
	rawURLS := []string{"https://example.com/path/to/resource", "https://test.io/another/path"}
	fuzzIdentifier := "FUZZ"
	fuzzableUrls, errs := GenerateFuzzableURLs(rawURLS, fuzzIdentifier)

	for _, u := range fuzzableUrls {
		t.Log(u)
	}

	if errs != nil {
		t.Logf("Error occured while generating fuzzable URLs: %v", errs)
	}

}
