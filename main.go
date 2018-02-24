package main

import (
	"fmt"
	"os"

	flags "github.com/jessevdk/go-flags"
)

// Version is filled in from git tags by the build script.
var Version = "undefined"

// O for options.
var O struct {
	Version    bool    `short:"v" description:"Display the version and exit."`
	Path       bool    `short:"p" description:"Strip the path from the filename, if present."`
	Suffix     bool    `short:"s" description:"Strip the suffix from the filename, if present."`
	OnlySuffix bool    `short:"S" description:"Return only the suffix."`
	Name       bool    `short:"n" description:"Strip the filename and suffix, leaving only the path."`
	Absolute   bool    `short:"a" description:"Return the absolute path."`
	File       FileOpt `positional-args:"true"`
}

// FileOpt is required.
type FileOpt struct {
	Name string `positional-arg-name:"FILE"`
}

func main() {
	_, err := flags.Parse(&O)
	if err != nil {
		pr("%$v", err)
		return
	}

	if O.Version {
		pr("name %s", Version)
		return
	}

	if O.File.Name == "" {
		epr("A filename is required.")
		return
	}

	parse(O.File.Name)
}

func pr(f string, v ...interface{}) {
	fmt.Printf(f+"\n", v...)
}

func epr(f string, v ...interface{}) {
	fmt.Fprintf(os.Stderr, f+"\n", v...)
}
