package sfz

import (
	"testing"
)

func TestRun(t *testing.T) {
	// Set up test data
	rawURLs := []string{"http://example.com"}
	fuzzIdentifier := "FUZZ"
	wordlist := "/path/to/wordlist.txt"
	outputJSON := "/path/to/output.json"
	outputFolder := "/path/to/output_folder"
	silent := false
	colorize := true
	headers := "User-Agent: Mozilla/5.0"
	disableAutoCalibration := false
	additionalFFUFArgs := []string{"-t", "50"}
	disableWarnings := false
	debugLog := false

	// Create the wrapper with the necessary options
	wrapper := New(
		WithRawURLs(rawURLs),
		WithFuzzIdentifier(fuzzIdentifier),
		WithWordlist(wordlist),
		WithFinalJSONOutput(outputJSON),
		WithFFUFResultsOutputFolder(outputFolder),
		WithSilentMode(silent),
		WithDisableColorizeOutput(!colorize),
		WithHeaders(headers),
		WithDisableAutomaticCalibration(disableAutoCalibration),
		WithAdditionalFFUFArgs(additionalFFUFArgs),
		WithDisableWarnings(disableWarnings),
		WithDebugLog(debugLog),
	)

	// Run the wrapper's Run method
	err := wrapper.Run()
	if err != nil {
		t.Errorf("Run() failed with error: %v", err)
	}
}
