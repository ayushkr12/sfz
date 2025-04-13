package sfz

import (
	"github.com/ayushkr12/sfz/internal/pkg/ffwrapper"
	"github.com/ayushkr12/sfz/internal/pkg/urlparser"
)

func RunFUFF(
	fuzzableUrls []string,
	wordlistPath string,
	outputFilePath string,
	headers string,
	disableColorizeOutput bool,
	disableAutomaticCalibration bool,
	silent bool,
	additionalFFUFArgs []string,
	debugLog bool,
) {
	ffufWrapper := ffwrapper.NewFFUFWrapper(
		fuzzableUrls, wordlistPath, outputFilePath, headers,
		disableAutomaticCalibration, disableColorizeOutput, silent, additionalFFUFArgs, debugLog,
	)
	ffufWrapper.LaunchCMDs()
}

func GenerateFuzzableURLs(rawURLs []string, fuzzIdentifier string) (fuzzableURLS []string, errs []error) {
	ctx := urlparser.NewFuzzContext(fuzzIdentifier)
	parsedURLs, err := ctx.UpdateURLsWithFuzzPaths(rawURLs)
	if err != nil {
		return nil, err
	}
	for _, parsedURL := range parsedURLs {
		for _, fuzzablePath := range parsedURL.FuzzablePaths {
			fuzzableURLS = append(fuzzableURLS, parsedURL.BaseURL+fuzzablePath)
		}
	}
	return fuzzableURLS, nil
}
