package ffwrapper

import (
	"fmt"
)

func (fw *FFUFWrapper) ValidateConfig() error {
	if fw.WordlistPath == "" {
		return fmt.Errorf("wordlist path is required")
	}
	return nil
}
