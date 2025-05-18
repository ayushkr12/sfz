# sfz

Smart fuzzin' cuz ffuf was boring

### What's this?

So I got tired of writing the same dumb ffuf URLs again and again like a caveman.
Wrote this tool for *me*, but hey, maybe you'll like it too.

You give it a bunch of URLs, it figures out where to slap that juicy `FUZZ`, then runs ffuf for you.
Don’t wanna give a wordlist? Cool, it’ll rip some paths & query params from the URLs and make a quick & dirty one on the fly.

Example:

```txt
http://localhost/api/core
```

It turns into stuff like:

* `http://localhost/api/FUZZ`
* `http://localhost/FUZZ/api`
* `http://localhost/api/core/FUZZ`

...then ffuf goes brrr.

### Usage

Run it like:

```bash
sfz -l urls.txt -w wordlist.txt
```

or just:

```bash
cat urls.txt | sfz
```

if you're lazy and want auto wordlist gen (grabs paths and query params from the input URLs using [xurl](https://github.com/ayushkr12/xurl) and cooks up a quick dirty wordlist by itself.).

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
