package cmd

var HelpMessage = `
sfz - Smart Fuzzing with ffuf

USAGE:
  sfz [flags]

FLAGS:

INPUT:
  -u,  --url                      Target URL to fuzz
  -l,  --list                     File containing multiple URLs to fuzz
  -i,  --fzi                      Placeholder string in URLs to replace with wordlist entries (default: "FUZZ")
  -w,  --wordlist                 Path to the wordlist file used for fuzzing

OUTPUT:
  -o,   --output-json             File path to store the final merged JSON output from all ffuf results
  -of,  --output-folder           Directory to store individual ffuf output files
  -kof, --keep-output-folder      Preserve the output folder after execution (default: false)

CUSTOMIZATION:
  -dfz, --disable-fuzz            Skip fuzzing and only generate fuzzable URLs
  -afa, --additional-ffuf-args    Extra arguments to pass directly to ffuf

FFUF OPTIONS:
  -H,   --headers                 Custom headers to include in requests
  -c,  --colorize                 Enable colored output for ffuf
  -dac, --disable-auto-calibration  Disable ffuf's automatic calibration feature

DEBUG:
  -d,  --debug-log                Enable debug mode to log detailed information
  -dw, --disable-warnings         Suppress warning messages

GENERAL:
  -s,  --silent                   Enable silent mode (minimal output)
  -h,  --help                     Show this help message
`
