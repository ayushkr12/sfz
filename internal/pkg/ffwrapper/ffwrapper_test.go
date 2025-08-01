package ffwrapper

import (
	"testing"
)

var fw = &FFUFWrapper{FuzzableURLs: []string{"http://example.com/FUZZ"}, Wordlist: []string{"test", "123"}, FFUFResultsOutputFolder: ".", Silent: false}

func TestLaunchCMDs(t *testing.T) {
	fw.LaunchCMDs()
}

func TestLaunchCMD(t *testing.T) {
	// Test cases for LaunchCMD function
	_, err := fw.LaunchCMD("https://example.com/FUZZ", ".")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
