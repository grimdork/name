# name

[![CI](https://github.com/grimdork/name/actions/workflows/ci.yml/badge.svg)](https://github.com/grimdork/name/actions/workflows/ci.yml)

Small, focused CLI that merges the behavior of `dirname` and `basename` with a few convenient extras. Designed for shell scripts and Makefiles where a compact, predictable filename-manipulation utility is useful.

## Install

- From source (recommended):

```sh
go install github.com/grimdork/name@latest
```

- Build locally:

```sh
make build
```

## Usage

```
name [OPTIONS]... [FILENAME]
```

- `name -h` shows help
- `name -v` prints the build version and build date (when provided via ldflags)

## Flags

- `-p`, `--nopath` — return only the base filename (strip directory path)
- `-s`, `--stripsuffix` — strip the suffix (file extension) from the name
- `-S`, `--onlysuffix` — return only the suffix (extension) without the leading dot
- `-n`, `--noname` — return only the path (directory component)
- `-a`, `--absolute` — expand the name to an absolute path

## Examples

```sh
# basename-like
name -p /usr/bin/env
# => env

# strip suffix
name -s main.go
# => main

# suffix only
name -S archive.tar.gz
# => gz

# dirname-like
name -n /sbin/ping6
# => /sbin

# absolute path
name -a main.go
# => /home/alice/src/name/main.go

# combined: absolute + strip suffix
name -as main.go
# => /home/alice/src/name/main
```

## Behavior notes and edge cases

- The `-S` (only suffix) flag takes precedence and returns immediately; other flags are not applied when `-S` is used.
- When `-n` (directory-only) is used on a bare filename (e.g. `foo`), `filepath.Dir("foo")` returns `.`; this tool normalises that to an empty string for scripting convenience.
- Suffix handling follows `filepath.Ext`: the suffix is the final `.` component. For `archive.tar.gz` the suffix is `gz` and `-s` removes only the final `.gz`.

## Development

- Run tests:

```sh
make test
```

- Vulnerability scan:

```sh
make govulncheck
```

- Linting is currently disabled in CI due to a runner/tooling compatibility issue; the Makefile includes a placeholder `lint` target.

## Contributing

Open issues and PRs are welcome. Please run the test suite and include a short description of your change.

## License

MIT — see LICENSE
