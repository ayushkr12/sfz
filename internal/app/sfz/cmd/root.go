package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	log "github.com/ayushkr12/sfz/pkg/logger"
	"github.com/ayushkr12/sfz/pkg/sfz"
	"github.com/ayushkr12/sfz/pkg/urlparser"
	"github.com/urfave/cli/v2"
	"golang.org/x/term"
)

func App() *cli.App {
	return &cli.App{
		Name:  "sfz",
		Usage: "Smart Fuzz using ffuf",
		Flags: Flags(),
		Action: func(c *cli.Context) error {
			return runMain()
		},
		CustomAppHelpTemplate: HelpMessage,
	}
}

func runMain() error {
	PrintBanner()

	var urls []string

	if debugLog {
		log.DisableDebug = false
	}

	// Check if a single URL is provided
	if url != "" {
		urls = append(urls, strings.TrimSpace(url))
		// Check if stdin is not a terminal (i.e., input is being piped)
		// If stdin is not a terminal, read urls from stdin
	} else if !term.IsTerminal(int(os.Stdin.Fd())) {
		log.Info("Reading URLs from stdin")
		reader := bufio.NewReader(os.Stdin)
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Error(fmt.Sprintf("Error reading stdin: %v", err))
				break
			}
			line = strings.TrimSpace(line)
			if line != "" {
				urls = append(urls, line)
			}
		}
	} else if urlFile != "" {
		// log.Info(fmt.Sprintf("Reading URLs from file: %s", urlFile))
		var err error
		urls, err = readURLsFromFile(urlFile)
		if err != nil {
			log.Error(fmt.Sprintf("Failed to read URLs from file: %v", err))
			return nil
		}
	} else {
		log.Info("No input provided. Use --list or pipe URLs into stdin.")
		return nil
	}

	if disableFuzz {
		fuzzableURLs, errs := urlparser.GenerateFuzzableURLs(urls, fuzzIdentifier)
		if len(fuzzableURLs) == 0 {
			log.Warn("No fuzzable URLs generated. Please Check your input.")
			return nil
		} else {
			log.Info(fmt.Sprintf("Generated %d fuzzable URLs\n", len(fuzzableURLs)))
		}
		for _, u := range fuzzableURLs {
			fmt.Println(u)
		}

		if len(errs) > 0 {
			log.Warn("Errors encountered during generating fuzzable urls: ")
			log.Warn(MergeErrorsToString(errs))
		}

		return nil
	}

	// Run the main fuzzing logic
	wrapper := sfz.New(
		sfz.WithRawURLs(urls),
		sfz.WithFuzzIdentifier(fuzzIdentifier),
		sfz.WithWordlist(wordlist),
		sfz.WithFinalJSONOutput(outputJSON),
		sfz.WithFFUFResultsOutputFolder(outputFolder),
		sfz.WithSilentMode(silent),
		sfz.WithDisableColorizeOutput(!colorize),
		sfz.WithHeaders(headers),
		sfz.WithDisableAutomaticCalibration(disableAutoCalibration),
		sfz.WithAdditionalFFUFArgs(additionalFFUFArgs.Value()),
		sfz.WithDisableWarnings(disableWarnings),
		sfz.WithDebugLog(debugLog),
	)

	err := wrapper.Run()
	if err != nil {
		return err
	}

	if !keepOutputFolder {
		RemoveDir(outputFolder)
	}

	return nil
}
