package ffwrapper

import (
	"fmt"
)

func (fw *FFUFWrapper) ValidateConfig() error {
	if len(fw.Wordlist) == 0 {
		return fmt.Errorf("empty wordlist provided")
	}

	if fw.FFUFResultsOutputFolder == "" && fw.FinalJSONOutputFilePath != "" {
		return fmt.Errorf("FFUF results output folder is required when final JSON output file path is set")
	}

	return nil
}
