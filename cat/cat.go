// cat written in Golang
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var lineNumber int
var emptyPrevLine bool

var AFlag = flag.Bool("A", false, "equivalent to -vET")
var bFlag = flag.Bool("b", false, "number nonempty output lines, overrides -n")
var eFlag = flag.Bool("e", false, "equivalent to -vE")
var EFlag = flag.Bool("E", false, "display $ at end of each line")
var sFlag = flag.Bool("s", false, "suppress repeated empty output lines")
var nFlag = flag.Bool("n", false, "number all output lines")
var tFlag = flag.Bool("t", false, "equivalent to -vT")
var TFlag = flag.Bool("T", false, "display TAB characters as ^I")
var uFlag = flag.Bool("u", false, "(ignored)")
var vFlag = flag.Bool("v", false, "use ^ and M- notation, except for LFD and TAB")
var helpFlag = flag.Bool("help", false, "display this help and exit")
var versionFlag = flag.Bool("version", false, "output version information and exit")

var versionInfo = `cat (cat written in go) 0.0.1
Written by Du Shaoshuai.
Available at github.com/dushaoshuai/go-linux.`

func init() {
	flag.Parse()
	if *AFlag {
		*EFlag = true
		*TFlag = true
	}
	if *eFlag {
		*EFlag = true
	}
	if *tFlag {
		*TFlag = true
	}
}

func main() {
	if *versionFlag { // version information
		fmt.Println(versionInfo)
		return
	}
	if *helpFlag { // help information
		usage()
		return
	}

	// No file name is supplied.
	if len(flag.Args()) == 0 {
		scan(os.Stdin)
	}

	// Scan all files.
	for _, filename := range flag.Args() {
		// File name is "-".
		if filename == "-" {
			scan(os.Stdin)
			continue
		}
		// Open file.
		file := openFile(filename)
		defer file.Close()
		// Scan file.
		scan(file)
	}
}

// usage displays help information, used by -help option.
func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [OPTION]... [FILE]...\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "Concatenate FILE(s) to standard output.\n\n")
	fmt.Fprintf(os.Stderr, "With no FILE, or when FILE is -, read standard input.\n\n")
	flag.PrintDefaults()
	fmt.Fprintf(os.Stderr, "\n-n is the same as --n.\n")
	fmt.Fprintf(os.Stderr, "-- stops options parsing.\n")
	fmt.Fprintf(os.Stderr, "\nBug: -v option doesn't work.\n")
	fmt.Fprintf(os.Stderr, "Available at github.com/dushaoshuai/go-linux.\n")
}

// scan scans one file.
func scan(file *os.File) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Taking into consideration of -s.
		// Suppress repeated empty output lines if needed.
		if skipOneLine(line) {
			continue
		}
		// Print one line taking into consideration of -n, -b, -E and -T.
		printLine(line)
	}
	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}

// openFile returns *os.File.
func openFile(filename string) *os.File {
	file, err := os.Open(filename)
	if err != nil {
		log.Println(err)
		return nil
	}
	return file
}

// -s, --squeeze-blank
// Suppress repeated empty output lines.
// skipOneLine returns true if we need to skip current empty line.
func skipOneLine(line string) bool {
	if *sFlag {
		if emptyPrevLine {
			if len(line) == 0 {
				// No need to set emptyPrevLine
				// to true because it's already true.
				return true
			} else {
				emptyPrevLine = false
			}
		} else {
			if len(line) == 0 {
				emptyPrevLine = true
			}
		} // No need to set emptyPrevLine to
		// false because it's already false.
	}
	return false
}

// printLine prints one line taking into consideration of -n, -b and -E.
// Wrapper of numberNonblankLine(), numberAllLine() and printOneLine().
func printLine(line string) {
	// Number nonempty output lines.
	if *bFlag {
		numberNonblankLine(line)
		return // -b overrides -n
	}
	// Number all output lines.
	if *nFlag {
		numberAllLine(line)
		return
	}
	lineNumber++
	printOneLine(line)
}

// -b, --number-nonblank
// Number nonempty output lines, overrides -n.
// Used by printLine().
func numberNonblankLine(line string) {
	if len(line) != 0 {
		lineNumber++
		fmt.Printf("%6d  ", lineNumber) // need improvements
	}
	printOneLine(line)
}

// -n, --number
// Number all output lines.
// Used by printLine().
func numberAllLine(line string) {
	lineNumber++
	fmt.Printf("%6d  ", lineNumber) // need improvements
	printOneLine(line)
}

// -E, --show-ends : Display $ at end of each line.
// -T, --show-tabs : display TAB characters as ^I.
// Used by printLine(), numberNonblankLine() and numberAllLine().
func printOneLine(line string) {
	if *EFlag {
		line = line + `$`
	}
	if *TFlag {
		line = strings.ReplaceAll(line, "\t", `^I`)
	}
	fmt.Println(line)
}
