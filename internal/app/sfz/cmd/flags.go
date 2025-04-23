package cmd

import "github.com/spf13/cobra"

var (
	fuzzIdentifier         = "FUZZ"
	urlFile                string
	rawURLs                []string
	wordlist               string
	outputJSON             string
	outputFolder           string
	silent                 = false
	colorize               bool
	headers                string
	disableAutoCalibration bool
	additionalFFUFArgs     []string
	disableWarnings        bool
	debugLog               bool

	rootCmd = &cobra.Command{
		Use:   "sfz",
		Short: "Smart Fuzz using ffuf",
		Long:  "Takes a list of URLs and a wordlist and smartly fuzzes different parts of the URL using ffuf\n",
		Run:   runMain,
	}
)

func ConfigureFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&urlFile, "list", "l", "", "File containing URLs to fuzz")
	cmd.Flags().StringVarP(&fuzzIdentifier, "fzi", "i", fuzzIdentifier, "Fuzz identifier to replace in URLs. (default: FUZZ)")
	cmd.Flags().StringVar(&wordlist, "wordlist", "w", "Path to wordlist")
	cmd.Flags().StringVar(&outputJSON, "output-json", "o", "Path to output the results as JSON")
	cmd.Flags().StringVar(&outputFolder, "output-folder", "of", "Path to output folder for FFUF results")
	cmd.Flags().BoolVar(&silent, "silent", silent, "Enable silent mode")
	cmd.Flags().BoolVar(&colorize, "colorize", true, "Enable or disable colorized output")
	cmd.Flags().StringVar(&headers, "headers", "", "Custom headers to send with requests")
	cmd.Flags().BoolVar(&disableAutoCalibration, "disable-auto-calibration", false, "Disable automatic calibration")
	cmd.Flags().BoolVar(&disableWarnings, "disable-warnings", false, "Disable warnings")
	cmd.Flags().BoolVar(&debugLog, "debug-log", false, "Enable debug logging")
	cmd.Flags().StringSliceVar(&additionalFFUFArgs, "additional-ffuf-args", nil, "Additional FFUF arguments")
}
