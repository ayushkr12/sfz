package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// parses a shell-style string like `-mc 200 -mr 'something'` into a []string
func ParseAdditionalFUFFArgs(input string) ([]string, error) {
	var args []string
	var current strings.Builder
	inQuote := false
	var quoteChar rune

	for _, r := range input {
		switch {
		case unicode.IsSpace(r) && !inQuote:
			if current.Len() > 0 {
				args = append(args, current.String())
				current.Reset()
			}
		case r == '\'' || r == '"':
			if inQuote {
				if r == quoteChar {
					inQuote = false
				} else {
					current.WriteRune(r)
				}
			} else {
				inQuote = true
				quoteChar = r
			}
		default:
			current.WriteRune(r)
		}
	}

	if current.Len() > 0 {
		args = append(args, current.String())
	}

	if inQuote {
		return nil, fmt.Errorf("unclosed quote in input: %s", input)
	}

	return args, nil
}

func MergeErrorsToString(errors []error) string {
	var sb strings.Builder
	for _, err := range errors {
		sb.WriteString(err.Error() + "\n")
	}
	return sb.String()
}

func readURLsFromFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
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

func RemoveDir(dirPath string) error {
	return os.RemoveAll(dirPath)
}
