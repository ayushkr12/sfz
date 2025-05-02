package cmd

import "github.com/urfave/cli/v2"

var (
	fuzzIdentifier         = "FUZZ"
	urlFile                string
	wordlist               string
	outputJSON             string
	outputFolder           string
	silent                 = false
	colorize               = true
	headers                string
	disableAutoCalibration = false
	additionalFFUFArgs     cli.StringSlice
	disableWarnings        = false
	debugLog               = false
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
			Usage:       `Fuzz identifier to replace in URLs (default: "FUZZ")`,
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
			Aliases:     []string{"s"},
			Usage:       "Enable silent mode (default: false)",
			Destination: &silent,
		},
		&cli.BoolFlag{
			Name:        "colorize",
			Aliases:     []string{"c"},
			Usage:       "Enable or disable colorized output (default: true)",
			Destination: &colorize,
		},
		&cli.StringFlag{
			Name:        "headers",
			Aliases:     []string{"H"},
			Usage:       "Custom headers to send with requests",
			Destination: &headers,
		},
		&cli.BoolFlag{
			Name:        "disable-auto-calibration",
			Aliases:     []string{"dac"},
			Usage:       "Disable automatic calibration (default: false)",
			Destination: &disableAutoCalibration,
		},
		&cli.StringSliceFlag{
			Name:        "additional-ffuf-args",
			Aliases:     []string{"afa"},
			Usage:       "Additional FFUF arguments",
			Destination: &additionalFFUFArgs,
		},
		&cli.BoolFlag{
			Name:        "disable-warnings",
			Aliases:     []string{"dw"},
			Usage:       "Disable warnings (default: false)",
			Destination: &disableWarnings,
		},
		&cli.BoolFlag{
			Name:        "debug-log",
			Aliases:     []string{"d"},
			Usage:       "Enable debug logging (default: false)",
			Destination: &debugLog,
		},
	}
}
