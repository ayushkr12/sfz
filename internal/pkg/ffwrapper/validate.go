package ffwrapper

import (
	"fmt"
)

func (fw *FFUFWrapper) Validate() error {
	if fw.TargetURL == "" {
		return fmt.Errorf("TargetURL is required")
	}

	if fw.WordlistPath == "" {
		return fmt.Errorf("WordlistPath is required")
	}

	return nil
}
