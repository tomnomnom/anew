package anewlib

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// appendDisplayLine - Choose to add the line `line`, if required to a file `fn`
func appendDisplayLine(f io.WriteCloser, fn string, line string, lines map[string]bool, quietMode bool) {
	// add the line to the map so we don't get any duplicates from stdin
	lines[line] = true

	if !quietMode {
		fmt.Println(line)
	}
	if fn != "" {
		fmt.Fprintf(f, "%s\n", line)
	}
}

// Anew - A public function which will append new lines `nl` provided as input
// to file `fn`. If quietMode `quietMode` enabled, then do not print lines to
// stdout. If `readFromStdin` is set to true, then lines to append will be read
// from stdin instead of `nl`. Returns newlines that are found.
func Anew(nl []string, fn string, quietMode bool, readFromStdin bool) []string {
	lines := make(map[string]bool)

	var newlines []string
	var f io.WriteCloser

	if fn != "" {
		// read the whole file into a map if it exists
		r, err := os.Open(fn)
		if err == nil {
			sc := bufio.NewScanner(r)

			for sc.Scan() {
				lines[sc.Text()] = true
			}
			r.Close()
		}

		// re-open the file for appending new stuff
		f, err = os.OpenFile(fn, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to open file for writing: %s\n", err)
			return newlines
		}
		defer f.Close()
	}

	// read the lines, append and output them if they're new
	if readFromStdin {
		sc := bufio.NewScanner(os.Stdin)
		for sc.Scan() {
			l := sc.Text()
			if lines[l] {
				continue
			}
			newlines = append(newlines, l)
			appendDisplayLine(f, fn, l, lines, quietMode)
		}
	} else {

		for l := range lines {
			if lines[l] {
				continue
			}
			appendDisplayLine(f, fn, l, lines, quietMode)
			newlines = append(newlines, l)
		}

	}
	return newlines
}
