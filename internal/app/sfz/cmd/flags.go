package cmd

import "github.com/urfave/cli/v2"

var (
	// Defaults
	defaultFuzzIdentifier         = "FUZZ"
	defaultSilent                 = false
	defaultColorize               = true
	defaultDisableAutoCalibration = false
	defaultDisableWarnings        = false
	defaultDebugLog               = false

	// Variables to hold flag values
	fuzzIdentifier         string
	urlFile                string
	wordlist               string
	outputJSON             string
	outputFolder           string
	silent                 bool
	colorize               bool
	headers                string
	disableAutoCalibration bool
	additionalFFUFArgs     cli.StringSlice
	disableWarnings        bool
	debugLog               bool
)

func Flags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "list",
			Aliases:     []string{"l"},
			Usage:       "File containing URLs to fuzz",
			Destination: &urlFile,
		},
		&cli.StringFlag{
			Name:        "fzi",
			Aliases:     []string{"i"},
			Usage:       "Fuzz identifier to replace in URLs",
			Value:       defaultFuzzIdentifier,
			Destination: &fuzzIdentifier,
		},
		&cli.StringFlag{
			Name:        "wordlist",
			Aliases:     []string{"w"},
			Usage:       "Path to wordlist",
			Destination: &wordlist,
		},
		&cli.StringFlag{
			Name:        "output-json",
			Aliases:     []string{"o"},
			Usage:       "Path to output the results as JSON",
			Destination: &outputJSON,
		},
		&cli.StringFlag{
			Name:        "output-folder",
			Aliases:     []string{"of"},
			Usage:       "Path to output folder for FFUF results",
			Destination: &outputFolder,
		},
		&cli.BoolFlag{
			Name:        "silent",
			Usage:       "Enable silent mode",
			Value:       defaultSilent,
			Destination: &silent,
		},
		&cli.BoolFlag{
			Name:        "colorize",
			Usage:       "Enable or disable colorized output",
			Value:       defaultColorize,
			Destination: &colorize,
		},
		&cli.StringFlag{
			Name:        "headers",
			Usage:       "Custom headers to send with requests",
			Destination: &headers,
		},
		&cli.BoolFlag{
			Name:        "disable-auto-calibration",
			Usage:       "Disable automatic calibration",
			Value:       defaultDisableAutoCalibration,
			Destination: &disableAutoCalibration,
		},
		&cli.StringSliceFlag{
			Name:        "additional-ffuf-args",
			Usage:       "Additional FFUF arguments",
			Destination: &additionalFFUFArgs,
		},
		&cli.BoolFlag{
			Name:        "disable-warnings",
			Usage:       "Disable warnings",
			Value:       defaultDisableWarnings,
			Destination: &disableWarnings,
		},
		&cli.BoolFlag{
			Name:        "debug-log",
			Usage:       "Enable debug logging",
			Value:       defaultDebugLog,
			Destination: &debugLog,
		},
	}
}
