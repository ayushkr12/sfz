package cmd

import "github.com/spf13/cobra"

var (
	fuzzIdentifier = "FUZZ"
	urlFile        string
	rootCmd        = &cobra.Command{
		Use:   "urlfuzzer",
		Short: "A CLI tool to parse and prepare URLs for fuzzing",
		Run:   runMain,
	}
)

func ConfigureFlags(cmd *cobra.Command) {
	cmd.Flags().StringVarP(&urlFile, "list", "l", "", "File containing URLs to fuzz")
	cmd.Flags().StringVarP(&fuzzIdentifier, "fzi", "i", fuzzIdentifier, "Fuzz identifier to replace in URLs. (default: FUZZ)")
}
