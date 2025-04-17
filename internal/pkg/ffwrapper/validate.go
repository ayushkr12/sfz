package ffwrapper

import (
	"fmt"
)

func (fw *FFUFWrapper) ValidateConfig() error {
	if fw.WordlistPath == "" {
		return fmt.Errorf("wordlist path is required")
	}

	if fw.FFUFResultsOutputFolder == "" {
		return fmt.Errorf("FFUF results output folder is required")
	}

	return nil
}
