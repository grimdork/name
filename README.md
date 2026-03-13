# name

Small, focused tool that merges the behavior of dirname and basename with a few convenient extras. Designed for shell scripts and Makefiles where a compact, predictable filename-manipulation utility is useful.

[![CI](https://github.com/grimdork/name/actions/workflows/ci.yml/badge.svg)](https://github.com/grimdork/name/actions/workflows/ci.yml)

What it does
- Return parts of a filename useful in shell scripting: path, base name, suffix, absolute path, and combinations.

Install
- From source (recommended when you want the latest):

  go install github.com/grimdork/name@latest

- Or build locally:

  make build

Usage
- name -h shows the help message.
- name -v prints the build version (when built with goreleaser or using the provided ldflags) and the build date.

Flags
- -p, --nopath    : Return only the base filename (strip directory path)
- -s, --stripsuffix : Strip the suffix (file extension) from the name
- -S, --onlysuffix : Return only the suffix (extension) without the leading dot
- -n, --noname    : Return only the path (directory component)
- -a, --absolute  : Expand to an absolute path

Examples
```sh
# basename-like
$ name -p /usr/bin/env
env

# strip suffix
$ name -s main.go
main

# suffix only
$ name -S archive.tar.gz
gz

# dirname-like
$ name -n /sbin/ping6
/sbin

# absolute path
$ name -a main.go
/home/alice/src/name/main.go

# combined: absolute and strip suffix
$ name -as main.go
/home/alice/src/name/main
```

Behaviour details and edge cases
- The -S (only suffix) flag takes precedence and returns immediately; other flags are not applied when -S is used.
- When -n (directory-only) is used on a bare filename (e.g. "foo"), the result of filepath.Dir is ".". For scripting convenience this tool normalises that to an empty string.
- Suffix handling: the suffix returned/removed is whatever filepath.Ext defines (the final '.' component). For files like "archive.tar.gz" the suffix is "gz" and -s removes only the final ".gz".

Development
- Run tests:

  make test

- Run the vulnerability scanner:

  make govulncheck

- Linting is currently disabled in CI due to a runner/tooling compatibility issue. The Makefile includes a placeholder lint target; we will re-enable it once the toolchain stabilises.

Contributing
- Open issues and PRs are welcome. Please run the test suite and include a short description of your change.
- When adding behaviour, prefer small, well-tested changes; keep the CLI stable for scripting consumers.

License
- MIT — see LICENSE

Maintainer
- Grim Dork (github.com/grimdork)
