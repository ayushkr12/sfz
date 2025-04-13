package ffwrapper

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	log "github.com/sirupsen/logrus"
)

type FFUFWrapper struct {
	TargetURLs                  []string // list of target URLs to scan
	FinalJSONOutputFilePath     string   // path to the merged JSON output file to create after all FFUF runs
	FFUFResultsOutputFolder     string   // path to the folder where FFUF results will be stored
	WordlistPath                string   // path to the wordlist file
	Headers                     string   // HTTP headers to be used in the requests
	DisableAutomaticCalibration bool     // flag to disable automatic calibration "-ac"
	DisableColorizeOutput       bool     // flag to disable colorized output "-c"
	Silent                      bool     // flag to run FFUF in silent mode "-s"
	AdditionalFFUFArgs          []string // additional arguments to pass to FFUF
	DebugLog                    bool     // log with timestamps for debugging
}

func NewFFUFWrapper(
	targetURLs []string,
	finalJSONOuputFilePath string,
	ffufResultsOutputFolder string,
	wordlistPath string,
	headers string,
	disableAutomaticCalibration bool,
	disableColorizeOutput bool,
	silent bool,
	additionalFFUFArgs []string,
	debugLog bool,
) *FFUFWrapper {
	return &FFUFWrapper{
		TargetURLs:                  targetURLs,
		FinalJSONOutputFilePath:     finalJSONOuputFilePath,
		FFUFResultsOutputFolder:     ffufResultsOutputFolder,
		WordlistPath:                wordlistPath,
		Headers:                     headers,
		DisableAutomaticCalibration: disableAutomaticCalibration,
		DisableColorizeOutput:       disableColorizeOutput,
		Silent:                      silent,
		AdditionalFFUFArgs:          additionalFFUFArgs,
		DebugLog:                    debugLog,
	}
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

		if url == "" {
			log.Warn("Skipping empty URL")
			continue
		}

		outputPath, err := fw.LaunchCMD(url, fw.FFUFResultsOutputFolder)
		if err != nil {
			log.Warn(fmt.Sprintf("Failed to launch FFUF for URL %s: %v", url, err))
			continue
		}

		log.Info(fmt.Sprintf("FFUF output for URL %s saved to %s", url, outputPath))
		ffufOutputFilePaths = append(ffufOutputFilePaths, outputPath)
	}

	if fw.FinalJSONOutputFilePath != "" {
		log.Info(fmt.Sprintf("Merging JSON output files into %s", fw.FinalJSONOutputFilePath))
		if err := MergeFFUFJSONOutputs(ffufOutputFilePaths, fw.FinalJSONOutputFilePath); err != nil {
			log.Error(fmt.Sprintf("Failed to merge JSON files: %v", err))
		}
	}
}

func (fw *FFUFWrapper) LaunchCMD(
	targetURL string,
	outputFolderPath string) (
	JSONOuputFilePath string,
	err error,
) {
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
	if fw.Silent {
		args = append(args, "-s")
	}

	if fw.FFUFResultsOutputFolder != "" { // set output file only if output folder is set
		JSONOuputFilePath = filepath.Join(outputFolderPath, GenerateRandomString(20)+".json")
		args = append(args, "-o", JSONOuputFilePath)
	}

	cmd := exec.Command("ffuf", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("error running ffuf: %w", err)
	}

	return JSONOuputFilePath, nil
}
