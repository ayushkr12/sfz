package sfz

type config struct {
	rawURLs        []string
	fuzzIdentifier string
	wordlist       string
	outputJSON     string
	outputFolder   string
	silent         bool
	colorize       bool
	headers        string
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

func WithWordlist(path string) Option {
	return func(c *config) {
		c.wordlist = path
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

func WithSilentMode(silent bool) Option {
	return func(c *config) {
		c.silent = silent
	}
}

func WithDisableColorizeOutput(disable bool) Option {
	return func(c *config) {
		c.colorize = disable
	}
}

func WithHeaders(headers string) Option {
	return func(c *config) {
		c.headers = headers
	}
}
