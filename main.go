package main

import (
	"flag"

	"github.com/manasmbellani/anew/anewlib"
)

func main() {
	var quietMode bool
	flag.BoolVar(&quietMode, "q", false, "quiet mode (no output at all)")
	flag.Parse()

	fn := flag.Arg(0)

	anewlib.Anew(nil, fn, quietMode, true)
}
