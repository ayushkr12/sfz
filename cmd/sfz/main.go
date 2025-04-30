package main

import (
	"os"

	"github.com/ayushkr12/sfz/internal/app/sfz/cmd"
	log "github.com/ayushkr12/sfz/pkg/logger"
)

func main() {
	if err := cmd.App().Run(os.Args); err != nil {
		log.Error(err.Error())
	}
}
