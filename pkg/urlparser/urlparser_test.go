package urlparser

import (
	"testing"

	"github.com/ayushkr12/sfz/internal/pkg/ffwrapper"
)

var rawURLS = []string{"https://example.com/path/to/resource", "https://test.io/another/path"}

func TestRunFUFF(t *testing.T) {
	fuzzableUrls := []string{"https://example.com/FUZZ", "https://test.io/FUZZ"}
	finalJSONOutputFilePath := "output.json"
	ffufResultsOutputFolder := "."
	wordlistPath := "-"
	headers := ""
	disableAutomaticCalibration := false
	disableColorizeOutput := false
	silent := false
	additionalFFUFArgs := []string{"-mc", "200"}
	debugLog := false

	w := ffwrapper.NewFFUFWrapper(
		fuzzableUrls, finalJSONOutputFilePath, ffufResultsOutputFolder, wordlistPath, headers, disableAutomaticCalibration, disableColorizeOutput, silent, additionalFFUFArgs, debugLog,
	)

	w.LaunchCMDs()
}

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
