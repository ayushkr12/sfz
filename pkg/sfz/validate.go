package sfz

import (
	"fmt"
)

func (w *Wrapper) Validate() error {
	if len(w.cfg.rawURLs) == 0 {
		return fmt.Errorf("no URLs provided")
	}
	if w.cfg.wordlistPath == "" && !w.cfg.enableAutoWordlist {
		return fmt.Errorf("no wordlist provided and auto wordlist generation is disabled")
	}
	if w.cfg.wordlistPath != "" && w.cfg.enableAutoWordlist {
		return fmt.Errorf("both custom wordlist and auto wordlist generation are provided")
	}
	return nil
}
