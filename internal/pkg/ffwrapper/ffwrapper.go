package ffwrapper

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

type FFUFWrapper struct {
	TargetURLs                  []string
	JSONOutputFilePath          string // path to the JSON output file
	OutputFolder                string
	WordlistPath                string
	Headers                     string
	DisableAutomaticCalibration bool
	DisableColorizeOutput       bool
	DisableSilent               bool
	AdditionalFFUFArgs          []string
	DebugLog                    bool
}

func (fw *FFUFWrapper) LaunchCMDs() {
	log.SetFormatter(&log.TextFormatter{DisableTimestamp: !fw.DebugLog})

	// Validate shared config once
	if err := fw.ValidateConfig(); err != nil {
		log.Error(fmt.Sprintf("failed to validate ffuf config: %v", err))
		return
	}

	var ffufOutputFilePaths []string
	for _, url := range fw.TargetURLs {
		log.Info(fmt.Sprintf("Launching FFUF for URL %s", url))

		// validate individual URL if needed
		if url == "" {
			log.Warn("Skipping empty URL")
			continue
		}

		outputPath, err := fw.LaunchCMD(url, fw.OutputFolder)
		if err != nil {
			log.Warn(fmt.Sprintf("Failed to launch FFUF for URL %s: %v", url, err))
			continue
		}

		log.Info(fmt.Sprintf("FFUF output for URL %s saved to %s", url, outputPath))
		ffufOutputFilePaths = append(ffufOutputFilePaths, outputPath)
	}

	if fw.JSONOutputFilePath != "" {
		log.Info(fmt.Sprintf("Merging JSON output files into %s", fw.JSONOutputFilePath))
		if err := MergeJsonFiles(ffufOutputFilePaths, fw.JSONOutputFilePath); err != nil {
			log.Error(fmt.Sprintf("Failed to merge JSON files: %v", err))
		}
	}
}

func (fw *FFUFWrapper) LaunchCMD(targetURL, outputFolderPath string) (string, error) {
	var args []string

	// Add additional args first to allow overriding
	args = append(args, fw.AdditionalFFUFArgs...)

	args = append(args, "-u", targetURL, "-w", fw.WordlistPath)

	if !fw.DisableColorizeOutput {
		args = append(args, "-c")
	}
	if fw.Headers != "" {
		args = append(args, "-H", fw.Headers)
	}
	if !fw.DisableAutomaticCalibration {
		args = append(args, "-ac")
	}
	if !fw.DisableSilent {
		args = append(args, "-s")
	}

	outputFile := filepath.Join(outputFolderPath, GenerateRandomString(20)+".json")
	args = append(args, "-o", outputFile)

	cmd := exec.Command("ffuf", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("error running ffuf: %w", err)
	}

	return outputFile, nil
}
