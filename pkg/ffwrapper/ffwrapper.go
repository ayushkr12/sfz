package ffwrapper

import (
	"github.com/ayushkr12/sfz/internal/pkg/ffwrapper"
)

type Wrapper struct {
	f *ffwrapper.FFUFWrapper
}

// New creates a new Wrapper instance with the given options.
func New(opts ...Option) (*Wrapper, error) {
	f := &ffwrapper.FFUFWrapper{}
	for _, opt := range opts {
		opt(f)
	}

	if err := f.ValidateConfig(); err != nil {
		return nil, err
	}

	return &Wrapper{f: f}, nil
}

func (w *Wrapper) RunFFUF() {
	w.f.LaunchCMDs()
}
