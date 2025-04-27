package tests

import (
	"os"
	"testing"

	"github.com/ayushkr12/sfz/pkg/sfz"
)

func TestRun(t *testing.T) {
	// Set up test data
	rawURLs := []string{"http://localhost:5000/api/v1/user/profile"}
	fuzzIdentifier := "FUZZ"
	wordlist := "testdata/wordlist.txt"
	outputJSON := "testdata/final_output.json"
	outputFolder := "testdata/output"
	silent := true
	colorize := true
	headers := "User-Agent: Mozilla/5.0"
	disableAutoCalibration := false
	additionalFFUFArgs := []string{"-fc", "403"}
	disableWarnings := false
	debugLog := false

	// Create the wrapper with the necessary options
	wrapper := sfz.New(
		sfz.WithRawURLs(rawURLs),
		sfz.WithFuzzIdentifier(fuzzIdentifier),
		sfz.WithWordlist(wordlist),
		sfz.WithFinalJSONOutput(outputJSON),
		sfz.WithFFUFResultsOutputFolder(outputFolder),
		sfz.WithSilentMode(silent),
		sfz.WithDisableColorizeOutput(!colorize),
		sfz.WithHeaders(headers),
		sfz.WithDisableAutomaticCalibration(disableAutoCalibration),
		sfz.WithAdditionalFFUFArgs(additionalFFUFArgs),
		sfz.WithDisableWarnings(disableWarnings),
		sfz.WithDebugLog(debugLog),
	)

	// Run the wrapper's Run method
	err := wrapper.Run()
	if err != nil {
		t.Errorf("Run() failed with error: %v", err)
	}
	_ = os.RemoveAll(outputFolder)
}
