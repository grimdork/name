package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/grimdork/climate/arg"
)

// These variables are set by goreleaser when building with that or GitHub actions.
var version = "undefined"
var date = "undefined"

func main() {
	opt := arg.New("name")
	opt.SetDefaultHelp(true)
	opt.SetOption(arg.GroupDefault, "v", "version", "Display the version and exit.", false, false, arg.VarBool, nil)
	opt.SetOption(arg.GroupDefault, "p", "nopath", "Strip the path from the filename, if present.", false, false, arg.VarBool, nil)
	opt.SetOption(arg.GroupDefault, "s", "stripsuffix", "Strip the suffix from the filename, if present.", false, false, arg.VarBool, nil)
	opt.SetOption(arg.GroupDefault, "S", "onlysuffix", "Return only the suffix.", false, false, arg.VarBool, nil)
	opt.SetOption(arg.GroupDefault, "n", "noname", "Strip the filename and suffix, leaving only the path.", false, false, arg.VarBool, nil)
	opt.SetOption(arg.GroupDefault, "a", "absolute", "Return the absolute path.", false, false, arg.VarBool, nil)
	opt.SetPositional("FILENAME", "Filename to process.", "", true, arg.VarString)

	var err error
	args := os.Args[1:]
	err = opt.Parse(args)
	if err != nil {
		if err == arg.ErrNoArgs {
			opt.PrintHelp()
			return
		}

		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(2)
	}

	if opt.GetBool("v") {
		pr("name %s", version)
		return
	}

	name := opt.GetPosString("FILENAME")
	if name == "" {
		opt.PrintHelp()
		return
	}

	if opt.GetBool("S") {
		suffix := filepath.Ext(name)
		if len(suffix) > 0 {
			pr("%s", suffix[1:])
		}
		return
	}

	if opt.GetBool("a") {
		var err error
		name, err = filepath.Abs(name)
		if err != nil {
			epr("Error: %s", err.Error())
			os.Exit(2)
		}
	}

	if opt.GetBool("p") {
		name = filepath.Base(name)
	}

	if opt.GetBool("n") {
		name = filepath.Dir(name)
	}

	if opt.GetBool("S") {
		suffix := filepath.Ext(name)
		if len(suffix) > 0 {
			name = strings.TrimSuffix(name, suffix)
		}
	}

	pr("%s", name)

}

func pr(f string, v ...interface{}) {
	fmt.Printf(f+"\n", v...)
}

func epr(f string, v ...interface{}) {
	fmt.Fprintf(os.Stderr, f+"\n", v...)
}
