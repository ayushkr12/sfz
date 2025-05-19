# sfz

Smart fuzzin' cuz ffuf was boring

### Usage

To generate fuzzable URLs only. use the `-dfz` flag:

```console
> cat url.txt
http://localhost/api/1/core/user
http://localhost/api/2/core/users
http://localhost/api/admin/pannel

> cat url.txt | sfz -dfz

            ____
   _____   / __/  ____
  / ___/  / /_   /_  /
 (__  )  / __/    / /_
/____/  /_/      /___/

  v1, with <3 by @ayushkr12

INFO Reading URLs from stdin
INFO Generated 18 fuzzable URLs
http://localhost/FUZZ
http://localhost/api/1/core/user/FUZZ
http://localhost/FUZZ/1/core/user
http://localhost/api/FUZZ/core/user
http://localhost/api/1/FUZZ/user
http://localhost/api/1/core/FUZZ
http://localhost/api/1/FUZZ
http://localhost/api/FUZZ
http://localhost/api/2/FUZZ/users
http://localhost/FUZZ/2/core/users
http://localhost/api/FUZZ/core/users
http://localhost/api/2/core/FUZZ
http://localhost/api/2/FUZZ
http://localhost/api/2/core/users/FUZZ
http://localhost/api/admin/FUZZ
http://localhost/api/admin/pannel/FUZZ
http://localhost/FUZZ/admin/pannel
http://localhost/api/FUZZ/pannel
```

To fuzz the above urls, either feed it a file or pipe URLs in:

```bash
sfz -l urls.txt -w wordlist.txt
```

or

```bash
cat urls.txt | sfz
```

sfz will grab paths and query params from URLs using [xurl](https://github.com/ayushkr12/xurl) and create a quick dirty wordlist on its own.

To pass custom options to `ffuf`, Use the `-afa` flag (Additional FFUF Args):

```bash
sfz -l urls.txt -w wordlist.txt -dac -afa "-mc 200 -mr 'sensitive'"
```

This passes `-mc 200 -mr 'sensitive'` to ffuf, so you can match status codes, regex, use filters, headers, delay, etc. anything ffuf supports.

### Flags

here’s the whole help menu.

```console
> sfz -h

sfz - Smart Fuzzing with ffuf

USAGE:
  sfz [flags]

FLAGS:

INPUT:
  -u,  --url                      Target URL to fuzz
  -l,  --list                     File containing multiple URLs to fuzz
  -i,  --fzi                      Placeholder string in URLs to replace with wordlist entries (default: "FUZZ")
  -w,  --wordlist                 Path to the wordlist file used for fuzzing
  -aw, --auto-wordlist            Automatically generate a wordlist from the provided URLs (enabled as default if no wordlist is provided)

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
```

### Install

Using go

```go
go install "github.com/ayushkr12/sfz/cmd/sfz@latest"
```

Clone this repo and build it.

```bash
git clone https://github.com/ayushkr12/sfz.git
cd sfz
make
./sfz -h
```

### using it as a package

If you wanna plug sfz into your own Go tools, you can use it like this:

```go
package main

import (
	"log"

	"github.com/ayushkr12/sfz/pkg/sfz"
)

func main() {
	wrapper := sfz.New(
		sfz.WithRawURLs([]string{
			"http://localhost:5000/api/v1/user",
			"http://localhost:5000/api/v1/admin?role=FUZZ",
		}),

		// Fuzzing config
		sfz.WithFuzzIdentifier("FUZZ"),
		sfz.WithWordlist("payloads.txt"),          // leave empty if you want auto wordlist
		sfz.WithEnableAutoWordlist(false),         // gets auto-enabled if no wordlist is given

		// Output
		sfz.WithFinalJSONOutput("results/final.json"),
		sfz.WithFFUFResultsOutputFolder("results"),

		// ffuf tweaks
		sfz.WithHeaders("User-Agent: Mozilla/5.0"),
		sfz.WithDisableAutomaticCalibration(false), // -ac flag in ffuf
		sfz.WithDisableColorizeOutput(false),       // -c flag in ffuf
		sfz.WithSilentMode(true),                   // -s flag in ffuf
		sfz.WithAdditionalFFUFArgs([]string{"-fc", "403"}),

		// logging
		sfz.WithDisableWarnings(false),
		sfz.WithDebugLog(true),
	)

	if err := wrapper.Run(); err != nil {
		log.Fatalf("sfz failed: %v", err)
	}
}
```

### Contribute

Sure, PRs welcome.

## Credits
this whole thing runs on top of [ffuf](https://github.com/ffuf/ffuf). so yeah, all fuzz magic comes from there. I just wrapped it up 'cause I was lazy. respect to @joohoi.
