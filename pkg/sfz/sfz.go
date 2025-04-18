package sfz

import (
	"github.com/ayushkr12/sfz/pkg/ffwrapper"
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
	fuzzableURLs, errs := urlparser.GenerateFuzzableURLs(w.cfg.rawURLs, w.cfg.fuzzIdentifier)
	if errs != nil {
		return errs[0] // todo: don't stop
	}

	ffOpts := []ffwrapper.Option{
		ffwrapper.WithTargetURLs(fuzzableURLs),
		ffwrapper.WithWordlist(w.cfg.wordlist),
		ffwrapper.WithFinalJSONOutput(w.cfg.outputJSON),
		ffwrapper.WithSilentMode(w.cfg.silent),
		ffwrapper.WithFFUFResultsOutputFolder(w.cfg.outputFolder),
		ffwrapper.WithDisableColorizeOutput(w.cfg.colorize),
		ffwrapper.WithHeaders(w.cfg.headers),
	}

	fw, err := ffwrapper.New(ffOpts...)
	if err != nil {
		return err
	}

	fw.RunFFUF()
	return nil
}
