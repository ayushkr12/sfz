package ffwrapper_test

import (
	"testing"

	"github.com/ayushkr12/sfz/pkg/ffwrapper"
)

func TestRunFFUF(t *testing.T) {
	wrapper, err := ffwrapper.New(
		ffwrapper.WithFuzzableURLs([]string{
			"https://example.com/FUZZ",
		}),
		ffwrapper.WithWordlist([]string{"test", "123"}),
		ffwrapper.WithFinalJSONOutput("test_output.json"),
		ffwrapper.WithSilentMode(true),
		ffwrapper.WithFFUFResultsOutputFolder("test_output"),
		ffwrapper.WithDisableColorizeOutput(true),
	)
	if err != nil {
		t.Fatalf("failed to create ffwrapper: %v", err)
	}

	wrapper.RunFFUF()
}
