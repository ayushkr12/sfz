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

	// Check if stdin is not a terminal (i.e., input is being piped)
	// If stdin is not a terminal, read urls from stdin
	if !term.IsTerminal(int(os.Stdin.Fd())) {
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
		log.Info(fmt.Sprintf("Reading URLs from file: %s", urlFile))
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

	fuzzableURLs, errs := urlparser.GenerateFuzzableURLs(urls, fuzzIdentifier)
	for _, u := range fuzzableURLs {
		fmt.Println(u)
	}

	if len(errs) > 0 {
		log.Warn("Errors encountered during fuzzing:")
		for _, e := range errs {
			log.Warn(e.Error())
		}
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

	return wrapper.Run()
}

func readURLsFromFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var urls []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			urls = append(urls, line)
		}
	}
	return urls, scanner.Err()
}
