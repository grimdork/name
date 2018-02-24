package main

import (
	"fmt"

	flags "github.com/jessevdk/go-flags"
)

// O for options.
var O struct {
	Path       bool    `short:"p" description:"Strip the path from the filename, if present."`
	Suffix     bool    `short:"s" description:"Strip the suffix from the filename, if present."`
	OnlySuffix bool    `short:"S" description:"Return only the suffix."`
	Name       bool    `short:"n" description:"Strip the filename and suffix, leaving only the path."`
	Absolute   bool    `short:"a" description:"Return the absolute path."`
	File       FileOpt `positional-args:"true"`
}

// FileOpt is required.
type FileOpt struct {
	Name string `positional-arg-name:"FILE" required:"true"`
}

func main() {
	_, err := flags.Parse(&O)
	if err != nil {
		return
	}

	parse(O.File.Name)
}

func pr(f string, v ...interface{}) {
	fmt.Printf(f+"\n", v...)
}
