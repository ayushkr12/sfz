package ffwrapper

import "github.com/ayushkr12/sfz/internal/pkg/ffwrapper"

type Option func(*ffwrapper.FFUFWrapper)

func WithFuzzableURLs(urls []string) Option {
	return func(f *ffwrapper.FFUFWrapper) {
		f.TargetURLs = urls
	}
}

func WithFinalJSONOutput(path string) Option {
	return func(f *ffwrapper.FFUFWrapper) {
		f.FinalJSONOutputFilePath = path
	}
}

func WithFFUFResultsOutputFolder(path string) Option {
	return func(f *ffwrapper.FFUFWrapper) {
		f.FFUFResultsOutputFolder = path
	}
}

func WithWordlist(path string) Option {
	return func(f *ffwrapper.FFUFWrapper) {
		f.WordlistPath = path
	}
}

func WithHeaders(headers string) Option {
	return func(f *ffwrapper.FFUFWrapper) {
		f.Headers = headers
	}
}

func WithDisableAutomaticCalibration(disable bool) Option {
	return func(f *ffwrapper.FFUFWrapper) {
		f.DisableAutomaticCalibration = disable
	}
}

func WithDisableColorizeOutput(disable bool) Option {
	return func(f *ffwrapper.FFUFWrapper) {
		f.DisableColorizeOutput = disable
	}
}

func WithSilentMode(silent bool) Option {
	return func(f *ffwrapper.FFUFWrapper) {
		f.Silent = silent
	}
}

func WithAdditionalFFUFArgs(args []string) Option {
	return func(f *ffwrapper.FFUFWrapper) {
		f.AdditionalFFUFArgs = args
	}
}

func WithDebugLog(enable bool) Option {
	return func(f *ffwrapper.FFUFWrapper) {
		f.DebugLog = enable
	}
}
