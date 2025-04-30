package cmd

import pkg "github.com/ayushkr12/sfz/pkg/version"

var Banner = `
            ____      
   _____   / __/  ____
  / ___/  / /_   /_  /
 (__  )  / __/    / /_
/____/  /_/      /___/
`

func PrintBanner() {
	println(Banner)
	println("  v" + pkg.Version + ", with <3 by @ayushkr12\n")
}
