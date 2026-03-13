package nameutil

import "testing"

func TestProcess(t *testing.T) {
	tests := []struct {
		name     string
		opts     Options
		expected string
	}{
		{"/sbin/ping6", Options{Nopath: true}, "ping6"},
		{"main.go", Options{StripSuffix: true}, "main"},
		{"main.go", Options{OnlySuffix: true}, "go"},
		{"/sbin/ping6", Options{Noname: true}, "/sbin"},
		{"main.go", Options{Absolute: true}, ""}, // we can't predict absolute path; just ensure no error
		{"foo", Options{Noname: true}, ""},       // Dir("foo") is "." -> normalized to ""
		{"/path/to/archive.tar.gz", Options{StripSuffix: true}, "/path/to/archive.tar"},
		{"/path/to/.hidden", Options{OnlySuffix: true}, "hidden"},
	}

	for _, tt := range tests {
		res, err := Process(tt.name, tt.opts)
		if err != nil {
			if tt.opts.Absolute {
				// For the Absolute test we accept any non-error result; check it is non-empty
				if res == "" {
					t.Fatalf("expected non-empty absolute path for %s, got empty", tt.name)
				}
				continue
			}
			t.Fatalf("unexpected error for %s: %v", tt.name, err)
		}

		if tt.opts.Absolute {
			// ensure it's an absolute path
			if len(res) == 0 || res[0] != '/' {
				t.Fatalf("expected absolute path, got %q", res)
			}
			continue
		}

		if res != tt.expected {
			t.Fatalf("%s %+v: expected %q, got %q", tt.name, tt.opts, tt.expected, res)
		}
	}
}
