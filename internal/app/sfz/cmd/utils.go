package cmd

import (
	"bufio"
	"os"
	"strings"
)

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
