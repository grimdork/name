package main

import (
	"fmt"
	"os"

	"github.com/Urethramancer/signor/opt"
)

// Version is filled in from git tags by the build script.
var version = "undefined"

// O for options.
var O struct {
	opt.DefaultHelp
	Version    bool   `short:"v" long:"version" help:"Display the version and exit."`
	Path       bool   `short:"p" long:"path" help:"Strip the path from the filename, if present."`
	Suffix     bool   `short:"s" long:"stripsuffix" help:"Strip the suffix from the filename, if present."`
	OnlySuffix bool   `short:"S" long:"onlysuffix" help:"Return only the suffix."`
	Name       bool   `short:"n" long:"noname" help:"Strip the filename and suffix, leaving only the path."`
	Absolute   bool   `short:"a" long:"absolute" help:"Return the absolute path."`
	File       string `placeholder:"FILENAME" help:"Filename to process."`
}

func main() {
	a := opt.Parse(&O)

	if O.Version {
		pr("name %s", version)
		return
	}

	if O.Help || O.File == "" {
		a.Usage()
		return
	}

	parse(O.File)
}

func pr(f string, v ...interface{}) {
	fmt.Printf(f+"\n", v...)
}

func epr(f string, v ...interface{}) {
	fmt.Fprintf(os.Stderr, f+"\n", v...)
}
