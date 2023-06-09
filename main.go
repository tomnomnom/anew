package main

import (
        "bufio"
        "flag"
        "fmt"
        "io"
        "os"
        "strings"
)

func main() {
        var quietMode bool
        var dryRun bool
        var trim bool
        var outputFile string

        flag.BoolVar(&quietMode, "q", false, "quiet mode (no output at all)")
        flag.BoolVar(&dryRun, "d", false, "don't append anything to the file, just print the new lines to stdout")
        flag.BoolVar(&trim, "t", false, "trim leading and trailing whitespace before comparison")
        flag.StringVar(&outputFile, "o", "", "specify the output file to write to if the line does not exist in any of the input files")
        flag.Parse()

        // Initialize a map for each file.
        lines := make(map[string]map[string]bool)
        filenames := flag.Args()
        var outputWriter io.WriteCloser

        if outputFile != "" && !dryRun {
                var err error
                outputWriter, err = os.OpenFile(outputFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
                if err != nil {
                        fmt.Fprintf(os.Stderr, "failed to open output file for writing: %s\n", err)
                        return
                }
                defer outputWriter.Close()
        }

        for _, fn := range filenames {
                lines[fn] = make(map[string]bool)

                // read the whole file into a map if it exists
                r, err := os.Open(fn)
                if err == nil {
                        sc := bufio.NewScanner(r)

                        for sc.Scan() {
                                if trim {
                                        lines[fn][strings.TrimSpace(sc.Text())] = true
                                } else {
                                        lines[fn][sc.Text()] = true
                                }
                        }
                        r.Close()
                }
        }

        // read the lines, append and output them if they're new
        sc := bufio.NewScanner(os.Stdin)

        for sc.Scan() {
                line := sc.Text()
                if trim {
                        line = strings.TrimSpace(line)
                }
                found := false
                for _, fn := range filenames {
                        if lines[fn][line] {
                                found = true
                                break
                        }
                }
                if found {
                        continue
                }

                // add the line to the map so we don't get any duplicates from stdin
                for _, fn := range filenames {
                        lines[fn][line] = true
                }

                if !quietMode {
                        fmt.Println(line)
                }
                if !dryRun && outputWriter != nil {
                        fmt.Fprintf(outputWriter, "%s\n", line)
                }
        }
}
