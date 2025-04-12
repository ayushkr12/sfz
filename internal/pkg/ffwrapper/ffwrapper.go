package ffwrapper

import (
	"os/exec"
)

type FFUFWrapper struct {
	TargetURL            string
	WordlistPath         string
	AutomaticCalibration bool
	AdditionalFFUFArgs   []string
}

func (fw *FFUFWrapper) LaunchCMD() *exec.Cmd {

	var args []string

	// Note: ffuf uses cobra for CLI which means the last argument wins
	// so we are Appending additional args first to allow for overriding default args.
	args = append(args, fw.AdditionalFFUFArgs...) // Append any additional FFUF arguments

	args = append(args,
		"-u", fw.TargetURL,
		"-w", fw.WordlistPath,
		"-json", "-s",
	)

	if fw.AutomaticCalibration {
		args = append(args, "-ac") // Automatically calibrate filtering options
	}

	cmd := exec.Command("ffuf", args...)

	return cmd

}
