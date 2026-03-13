package nameutil

import (
	"path/filepath"
	"strings"
)

// Options controls how a filename is processed.
type Options struct {
	Nopath      bool // -p
	StripSuffix bool // -s
	OnlySuffix  bool // -S
	Noname      bool // -n
	Absolute    bool // -a
}

// Process applies the options to the provided filename and returns the result.
// Returns an error only for filesystem operations (e.g. Abs failure).
func Process(name string, opts Options) (string, error) {
	if opts.OnlySuffix {
		suffix := filepath.Ext(name)
		if len(suffix) > 0 {
			return strings.TrimPrefix(suffix, "."), nil
		}
		return "", nil
	}

	var err error
	if opts.Absolute {
		name, err = filepath.Abs(name)
		if err != nil {
			return "", err
		}
	}

	if opts.Nopath {
		name = filepath.Base(name)
	}

	if opts.Noname {
		name = filepath.Dir(name)
		// Normalize '.' to empty string for bare filenames when -n is used.
		if name == "." {
			name = ""
		}
	}

	if opts.StripSuffix {
		suffix := filepath.Ext(name)
		if len(suffix) > 0 {
			name = strings.TrimSuffix(name, suffix)
		}
	}

	return name, nil
}
