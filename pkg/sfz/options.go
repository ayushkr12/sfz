package sfz

type config struct {
	rawURLs                []string
	fuzzIdentifier         string
	wordlist               string
	outputJSON             string
	outputFolder           string
	silent                 bool
	colorize               bool
	headers                string
	disableAutoCalibration bool
	additionalFFUFArgs     []string
	disableWarnings        bool
	debugLog               bool
}

type Option func(*config)

func WithRawURLs(urls []string) Option {
	return func(c *config) {
		c.rawURLs = urls
	}
}

func WithFuzzIdentifier(fuzz string) Option {
	return func(c *config) {
		c.fuzzIdentifier = fuzz
	}
}

func WithFinalJSONOutput(path string) Option {
	return func(c *config) {
		c.outputJSON = path
	}
}

func WithFFUFResultsOutputFolder(path string) Option {
	return func(c *config) {
		c.outputFolder = path
	}
}

func WithWordlist(path string) Option {
	return func(c *config) {
		c.wordlist = path
	}
}

func WithHeaders(headers string) Option {
	return func(c *config) {
		c.headers = headers
	}
}

func WithDisableAutomaticCalibration(disable bool) Option {
	return func(c *config) {
		c.disableAutoCalibration = disable
	}
}

func WithDisableColorizeOutput(disable bool) Option {
	return func(c *config) {
		c.colorize = !disable
	}
}

func WithSilentMode(silent bool) Option {
	return func(c *config) {
		c.silent = silent
	}
}

func WithAdditionalFFUFArgs(args []string) Option {
	return func(c *config) {
		c.additionalFFUFArgs = args
	}
}

func WithDisableWarnings(disable bool) Option {
	return func(c *config) {
		c.disableWarnings = disable
	}
}

func WithDebugLog(enable bool) Option {
	return func(c *config) {
		c.debugLog = enable
	}
}
