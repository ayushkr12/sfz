package ffwrapper

import (
	"fmt"
)

func (fw *FFUFWrapper) ValidateConfig() error {
	if fw.WordlistPath == "" {
		return fmt.Errorf("wordlist path is required")
	}
	if fw.OutputFolder == "" {
		return fmt.Errorf("output folder is required")
	}
	return nil
}
