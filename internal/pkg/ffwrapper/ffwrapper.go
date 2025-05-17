package ffwrapper

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"

	log "github.com/ayushkr12/sfz/pkg/logger"
)

type FFUFWrapper struct {
	FuzzableURLs                []string // list of target URLs to scan
	FinalJSONOutputFilePath     string   // path to the merged JSON output file to create after all FFUF runs
	FFUFResultsOutputFolder     string   // path to the folder where FFUF results will be stored
	Wordlist                    []string // list of words to use for fuzzing
	Headers                     string   // HTTP headers to be used in the requests
	DisableAutomaticCalibration bool     // flag to disable automatic calibration "-ac"
	DisableColorizeOutput       bool     // flag to disable colorized output "-c"
	Silent                      bool     // flag to run FFUF in silent mode "-s"
	AdditionalFFUFArgs          []string // additional arguments to pass to FFUF
	DebugLog                    bool     // log with timestamps for debugging
}

func (fw *FFUFWrapper) LaunchCMDs() {
	if !fw.DebugLog {
		log.DisableDebug = true
	}

	// Validate shared config once
	if err := fw.ValidateConfig(); err != nil {
		log.Error(fmt.Sprintf("failed to validate ffuf config: %v", err))
		return
	}

	var ffufOutputFilePaths []string
	for _, url := range fw.FuzzableURLs {
		if url == "" {
			log.Warn("Skipping empty URL")
			continue
		}

		outputPath, err := fw.LaunchCMD(url, fw.FFUFResultsOutputFolder)
		if err != nil {
			log.Warn(fmt.Sprintf("Failed to launch FFUF for URL %s: %v", url, err)) // todo: return those errors slice as well as printing them
			continue
		}
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

	args = append(args, "-u", targetURL, "-w", "-") // "-" means to read from stdin

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
		err = CreateFolderIfNotExists(outputFolderPath)
		if err != nil {
			return "", fmt.Errorf("failed to create output folder: %v", err)
		}
		JSONOuputFilePath = filepath.Join(outputFolderPath, GenerateRandomString(20)+".json")
		args = append(args, "-o", JSONOuputFilePath)
	}

	fmt.Println() // for better readability
	log.Info(fmt.Sprintf("Launching FFUF for URL %s", targetURL))
	// fmt.Println() // for better readability
	log.Debug(fmt.Sprintf("Executing FFUF command: %s\n", getRawCommandOutput(fw.AdditionalFFUFArgs)))

	pr, pw := io.Pipe()

	cmd := exec.Command("ffuf", args...)
	cmd.Stdin = pr
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	go func() { // write the wordlist to the pipe
		defer pw.Close()
		for _, word := range fw.Wordlist {
			if word == "" {
				continue
			}
			pw.Write([]byte(word + "\n"))
		}
	}()

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("error running ffuf: %w", err)
	}

	if outputFolderPath != "" {
		log.Info(
			fmt.Sprintf("Saved to %s\n", JSONOuputFilePath),
		)
	}

	return JSONOuputFilePath, nil
}
