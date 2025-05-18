package sfz

import (
	"fmt"

	"github.com/ayushkr12/sfz/pkg/ffwrapper"
	"github.com/ayushkr12/sfz/pkg/logger"
	"github.com/ayushkr12/sfz/pkg/urlparser"
	"github.com/ayushkr12/sfz/pkg/utils"
	xurl "github.com/ayushkr12/xurl/pkg"
)

type Wrapper struct {
	cfg *config
}

// New creates a new sfz.Wrapper with unified sfz options
func New(opts ...Option) *Wrapper {
	cfg := &config{}
	for _, opt := range opts {
		opt(cfg)
	}
	return &Wrapper{cfg: cfg}
}

// Run generates fuzzable URLs and launches ffuf with mapped options
func (w *Wrapper) Run() error {
	if !w.cfg.debugLog {
		logger.DisableDebug = true
	} else {
		logger.DisableDebug = false
	}

	if w.cfg.disableWarnings {
		logger.DisableWarn = true
	}

	var wordlist []string
	var err error

	if w.cfg.wordlistPath == "" {
		logger.Warn("No wordlist provided. Wordlist will be generated automatically.")
		w.cfg.enableAutoWordlist = true // Force auto wordlist generation if no wordlist is provided
	} else {
		wordlist, err = utils.FileToSlice(w.cfg.wordlistPath)
		if err != nil {
			return fmt.Errorf("failed to read wordlist file: %v", err)
		}
	}

	if w.cfg.enableAutoWordlist {
		logger.Info("Generating custom wordlist from input URLs")
		var errs []error
		wordlist, errs = xurl.ExtractWords(w.cfg.rawURLs)
		if len(wordlist) == 0 {
			return fmt.Errorf("no words found in the provided URLs")
		}
		logger.Info(fmt.Sprintf("Generated %d words from input URLs\n", len(wordlist)))
		logger.Debug(fmt.Sprintf("Generated words: %v", wordlist))
		if len(errs) > 0 {
			for _, err := range errs {
				logger.Warn(err.Error())
			}
		}
	}

	fuzzableURLs, errs := urlparser.GenerateFuzzableURLs(w.cfg.rawURLs, w.cfg.fuzzIdentifier)
	if len(errs) > 0 {
		for _, err := range errs {
			if !w.cfg.disableWarnings {
				logger.Warn(err.Error())
			}
		}
	}

	if len(fuzzableURLs) == 0 {
		logger.Info("No fuzzable URLs generated. Exiting.")
		return nil
	}

	logger.Info(fmt.Sprintf("Generated %d fuzzable URLs\n", len(fuzzableURLs)))

	ffOpts := []ffwrapper.Option{
		ffwrapper.WithFuzzableURLs(fuzzableURLs),
		ffwrapper.WithWordlist(wordlist),
		ffwrapper.WithFinalJSONOutput(w.cfg.outputJSON),
		ffwrapper.WithSilentMode(w.cfg.silent),
		ffwrapper.WithFFUFResultsOutputFolder(w.cfg.outputFolder),
		ffwrapper.WithDisableColorizeOutput(!w.cfg.colorize),
		ffwrapper.WithHeaders(w.cfg.headers),
		ffwrapper.WithDisableAutomaticCalibration(w.cfg.disableAutoCalibration),
		ffwrapper.WithAdditionalFFUFArgs(w.cfg.additionalFFUFArgs),
		ffwrapper.WithDebugLog(w.cfg.debugLog),
	}

	fw, err := ffwrapper.New(ffOpts...)
	if err != nil {
		return err
	}

	fw.RunFFUF()
	return nil
}
