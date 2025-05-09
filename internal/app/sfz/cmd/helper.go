package cmd

var HelpMessage = `
sfz - Smart Fuzz using ffuf

USAGE:
  sfz [flags]

FLAGS:

INPUT:
  -l, -list                     file containing URLs to fuzz
  -i, -fzi                      fuzz identifier to replace in URLs (default: "FUZZ")
  -w, -wordlist                 path to wordlist

OUTPUT:
  -o, -output-json              path to save results as JSON
  -of, -output-folder           path to output folder for ffuf results

CUSTOMIZATION:
  -H, -headers                  custom headers to send with requests
  -afa, -additional-ffuf-args   additional FFUF arguments

BEHAVIOR:
  -dfz, -disable-fuzz              disable fuzzing and generate Fuzzable URLs only
  -dac, -disable-auto-calibration  disable automatic calibration
  -dw, -disable-warnings           disable warnings

DEBUG:
  -d, -debug-log                enable debug logging
  -s, -silent                   enable silent mode
  -c, -colorize                 enable colorized output

MISC:
  -h, -help                     show help
`
