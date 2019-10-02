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

	lines := make(map[string]struct{})

	// Read the input to append to a map. This is done since the
	// data to be appended would be less than data to which
	// it is to be appended. This saves memory.
	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		line := sc.Text()
		if _, ok := lines[line]; !ok {
			lines[line] = struct{}{}
		}
	}

	// Re-Open the file in append mode. Read the lines and if any of them
	// match with the line in our map, remove it from the map.
	f, err := os.OpenFile(fn, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open file for writing: %s\n", err)
	}
	sc = bufio.NewScanner(f)

	for sc.Scan() {
		line := sc.Text()
		if _, ok := lines[line]; ok {
			delete(lines, line)
			continue
		}
	}

	// Finally write all the remaining elements in input map to the output
	// as they now contain only unique words.
	for line := range lines {
		fmt.Println(line)
		f.WriteString(line + "\n")
	}
	f.Close()
}
