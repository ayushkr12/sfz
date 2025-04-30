package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/ayushkr12/sfz/pkg/sfz"
	"github.com/ayushkr12/sfz/pkg/urlparser"
	log "github.com/sirupsen/logrus"
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
	}
}

func runMain() error {
	var urls []string

	if !term.IsTerminal(int(os.Stdin.Fd())) {
		log.Info("Reading URLs from stdin...")
		reader := bufio.NewReader(os.Stdin)
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Errorf("Error reading stdin: %v", err)
				break
			}
			line = strings.TrimSpace(line)
			if line != "" {
				urls = append(urls, line)
			}
		}
	} else if urlFile != "" {
		log.Infof("Reading URLs from file: %s", urlFile)
		var err error
		urls, err = readURLsFromFile(urlFile)
		if err != nil {
			log.Fatalf("Failed to read URLs from file: %v", err)
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
			log.Warnf("- %v", e)
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
