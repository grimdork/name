package main

import (
	"fmt"
	"os"

	"github.com/grimdork/climate/arg"
	"github.com/grimdork/name/nameutil"
)

// These variables are set by goreleaser when building with that or GitHub actions.
var version = "undefined"
var date = "undefined"

func main() {
	opt := arg.New("name","Manipulate path names.")
	opt.SetDefaultHelp(true)
	opt.SetOption(arg.GroupDefault, "v", "version", "Display the version and exit.", false, false, arg.VarBool, nil)
	opt.SetOption(arg.GroupDefault, "p", "nopath", "Strip the path from the filename, if present.", false, false, arg.VarBool, nil)
	opt.SetOption(arg.GroupDefault, "s", "stripsuffix", "Strip the suffix from the filename, if present.", false, false, arg.VarBool, nil)
	opt.SetOption(arg.GroupDefault, "S", "onlysuffix", "Return only the suffix.", false, false, arg.VarBool, nil)
	opt.SetOption(arg.GroupDefault, "n", "noname", "Strip the filename and suffix, leaving only the path.", false, false, arg.VarBool, nil)
	opt.SetOption(arg.GroupDefault, "a", "absolute", "Return the absolute path.", false, false, arg.VarBool, nil)
	opt.SetPositional("FILENAME", "Filename to process.", "", true, arg.VarString)

	args := os.Args[1:]
	if err := opt.Parse(args); err != nil {
		if err == arg.ErrNoArgs {
			opt.PrintHelp()
			return
		}

		fmt.Printf("Error: %s\n", err.Error())
		os.Exit(2)
	}

	if opt.GetBool("v") {
		pr("name %s (%s)", version, date)
		return
	}

	name := opt.GetPosString("FILENAME")
	if name == "" {
		opt.PrintHelp()
		return
	}

	opts := nameutil.Options{
		Nopath:      opt.GetBool("p"),
		StripSuffix: opt.GetBool("s"),
		OnlySuffix:  opt.GetBool("S"),
		Noname:      opt.GetBool("n"),
		Absolute:    opt.GetBool("a"),
	}

	res, err := nameutil.Process(name, opts)
	if err != nil {
		epr("Error: %s", err.Error())
		os.Exit(2)
	}

	pr("%s", res)
}

func pr(f string, v ...interface{}) {
	fmt.Printf(f+"\n", v...)
}

func epr(f string, v ...interface{}) {
	fmt.Fprintf(os.Stderr, f+"\n", v...)
}
