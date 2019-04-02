package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()

	fn := flag.Arg(0)
	if fn == "" {
		fmt.Println("usage: anew <file>")
		os.Exit(1)
	}

	lines := make(map[string]bool)

	// read the whole file into a map if it exists
	f, err := os.Open(fn)
	if err == nil {
		sc := bufio.NewScanner(f)

		for sc.Scan() {
			lines[sc.Text()] = true
		}
		f.Close()
	}

	// re-open the file for appending new stuff
	f, err = os.OpenFile(fn, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open file for writing: %s\n", err)
	}
	defer f.Close()

	// read the lines, append and output them if they're new
	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		line := sc.Text()
		if lines[line] {
			continue
		}

		// add the line to the map so we don't get any duplicates from stdin
		lines[line] = true

		fmt.Println(line)
		f.WriteString(line + "\n")
	}
}
