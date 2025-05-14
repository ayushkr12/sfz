package cmd

import "github.com/urfave/cli/v2"

var (
	fuzzIdentifier         = "FUZZ"
	urlFile                string
	url                    string
	disableFuzz            = false
	wordlist               string
	outputJSON             string
	outputFolder           string
	keepOutputFolder       = false
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
			Name:        "url",
			Aliases:     []string{"u"},
			Usage:       "Single URL to fuzz",
			Destination: &url,
		},
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
			Value:       fuzzIdentifier,
			Destination: &fuzzIdentifier,
		},
		&cli.BoolFlag{
			Name:        "disable-fuzz",
			Aliases:     []string{"dfz"},
			Usage:       `Disable fuzzing and generate Fuzzable URLs only (default: false)`,
			Value:       disableFuzz,
			Destination: &disableFuzz,
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
			Name:        "keep-output-folder",
			Aliases:     []string{"kof"},
			Usage:       "Disable deleting output folder after run (default: false)",
			Destination: &keepOutputFolder,
		},
		&cli.BoolFlag{
			Name:        "silent",
			Aliases:     []string{"s"},
			Usage:       "Run FUFF with silent mode enabled (default: false)",
			Value:       silent,
			Destination: &silent,
		},
		&cli.BoolFlag{
			Name:        "colorize",
			Aliases:     []string{"c"},
			Usage:       "Run FUFF with colorize mode enabled (default: true)",
			Value:       colorize,
			Destination: &colorize,
		},
		&cli.StringFlag{
			Name:        "headers",
			Aliases:     []string{"H"},
			Usage:       "Custom FUFF headers to send with requests",
			Destination: &headers,
		},
		&cli.BoolFlag{
			Name:        "disable-auto-calibration",
			Aliases:     []string{"dac"},
			Usage:       "Disable FUFF automatic calibration (default: false)",
			Value:       disableAutoCalibration,
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
			Value:       disableWarnings,
			Destination: &disableWarnings,
		},
		&cli.BoolFlag{
			Name:        "debug-log",
			Aliases:     []string{"d"},
			Usage:       "Enable debug logging (default: false)",
			Value:       debugLog,
			Destination: &debugLog,
		},
	}
}
