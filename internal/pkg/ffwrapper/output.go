package ffwrapper

import (
	"encoding/json"
	"fmt"
	"io"
	"maps"
	"os"
)

func MergeFFUFJSONOutputs(inputFilePaths []string, outputFilePath string) error {
	return mergeJsonFiles(inputFilePaths, outputFilePath)
}

func mergeJsonFiles(filePaths []string, outputFilePath string) error {
	merged := make(map[string]any)

	for _, path := range filePaths {
		file, err := os.Open(path)
		if err != nil {
			return fmt.Errorf("failed to open file %s: %w", path, err)
		}
		defer file.Close()

		data, err := io.ReadAll(file)
		if err != nil {
			return fmt.Errorf("failed to read file %s: %w", path, err)
		}

		var temp map[string]any
		if err := json.Unmarshal(data, &temp); err != nil {
			return fmt.Errorf("failed to parse JSON in file %s: %w", path, err)
		}

		maps.Copy(merged, temp)
	}

	outputData, err := json.MarshalIndent(merged, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal merged JSON: %w", err)
	}

	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer outputFile.Close()

	if _, err := outputFile.Write(outputData); err != nil {
		return fmt.Errorf("failed to write to output file: %w", err)
	}

	return nil
}
