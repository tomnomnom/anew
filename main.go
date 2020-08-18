package main

import (
	"flag"

	"./anew"
)

func main() {
	var quietMode bool
	flag.BoolVar(&quietMode, "q", false, "quiet mode (no output at all)")
	flag.Parse()

	fn := flag.Arg(0)

	anew.Anew(nil, fn, quietMode, true)
}
