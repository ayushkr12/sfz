package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"strings"

	"github.com/ayushkr12/sfz/pkg/sfz"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

func runMain(_ *cobra.Command, _ []string) {
	var rawURLs []string

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
				rawURLs = append(rawURLs, line)
			}
		}
	} else if urlFile != "" {
		log.Infof("Reading URLs from file: %s", urlFile)
		var err error
		rawURLs, err = readURLsFromFile(urlFile)
		if err != nil {
			log.Fatalf("Failed to read URLs from file: %v", err)
		}
	} else {
		log.Info("No input provided. Use --list or pipe URLs into stdin.")
	}

	fuzzableURLs, errs := sfz.GenerateFuzzableURLs(rawURLs, fuzzIdentifier)

	for _, u := range fuzzableURLs {
		fmt.Println(u)
	}

	if len(errs) > 0 {
		log.Warn("Errors encountered during fuzzing:")
		for _, e := range errs {
			log.Warnf("- %v", e)
		}
	}
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

func Execute() {
	ConfigureFlags(rootCmd)
	if err := rootCmd.Execute(); err != nil {
		log.Error(err)
	}
}
