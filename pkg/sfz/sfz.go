package sfz

import (
	"fmt"

	"github.com/ayushkr12/sfz/pkg/ffwrapper"
	"github.com/ayushkr12/sfz/pkg/logger"
	"github.com/ayushkr12/sfz/pkg/urlparser"
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
	}

	if w.cfg.wordlist == "" {
		logger.Error("Wordlist is required")
	}

	if w.cfg.disableWarnings {
		logger.DisableWarn = true
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
		ffwrapper.WithWordlist(w.cfg.wordlist),
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
