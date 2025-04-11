package urlparser

import (
	"testing"
)

func TestUpdateURLsWithFuzzPaths(t *testing.T) {
	ctx := &FuzzContext{FuzzIdentifier: "FUZZ"}
	rawURLs := []string{
		"https://www.example.com/path/to/resource",
		"https://www.example.com/another/path",
		"https://www.another-domain.com/some/path",
	}
	parsedURLs, err := ctx.UpdateURLsWithFuzzPaths(rawURLs)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	for _, p := range parsedURLs {
		t.Logf("BaseURL: %s, Paths: %v", p.BaseURL, p.FuzzablePaths)
	}
}

func TestUpdateURLWithFuzzPaths(t *testing.T) {
	ctx := &FuzzContext{FuzzIdentifier: "FUZZ"}

	testURL := "https://www.example.com/path/to/resource"
	parsedURLs, err := ctx.UpdateURLWithFuzzPaths(testURL)

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	for _, p := range parsedURLs {
		t.Logf("BaseURL: %s, Paths: %v", p.BaseURL, p.FuzzablePaths)
	}

}
