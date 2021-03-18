package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var nFlag = flag.Bool("n", false, "do not output the trailing newline")
var eFlag = flag.Bool("e", false, "enable interpretation of backslash escapes")
var EFlag = flag.Bool("E", true, "disable interpretation of backslash escapes")

var vFlag = flag.Bool("version", false, "output version information and exit")
var versionInfo = `ecgo (echo written in go) 0.0.1
Written by Du Shaoshuai.`

func main() {
	flag.Parse()

	// version information
	if *vFlag {
		fmt.Println(versionInfo)
		return
	}

	// backslash escapes

	// "-e" and "-E" may be used together, the latter
	// appeared in the command line overrides the former
	interpretation := false
	for _, s := range os.Args[1:] {
		if !strings.HasPrefix(s, "-") || s == "--" {
			break
		}
		if s == "-e" {
			interpretation = true
		} else if s == "-E" {
			interpretation = false
		}
	}

	args := strings.Join(flag.Args(), " ")
	if *eFlag && interpretation { // interprete
		args = strings.ReplaceAll(args, "\\\\", "\\")
		var err error
		args, err = strconv.Unquote(`"` + args + `"`)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Print(args)

	// trailing newline
	if !*nFlag {
		fmt.Println()
	}
}
