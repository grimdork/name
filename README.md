# name
Merges functionality of the dirname and basename shell commands with a few extra features.

## What
This program returns segments of a filename which are useful for shell scripting.

## Usage
`name -h` for help.

`name -v` to show version (applicable if built with the `build` script).

`name <filename>` returns the `filename` as is. Mostly useless.

`name -p <filename>` returns `filename` without any path or leading slashes. Example:
```sh
$ name -p /sbin/ping6
ping6
```

`name -s <filename>` returns `filename` without any suffix. Example:
```sh
$ name -s parse.go
parse
```

`name -S <filename>` returns only the suffix of `filename`. Example:
```sh
$ name -S parse.go
go
```

`name -n <filename>` returns only the path of `filename`. Example:
```sh
$ name -n /sbin/ping6
/sbin
```

`name -a <filename>` expands `filename` to include the full (absolute) path. Example:
```sh
$ name -a parse.go
/Users/orb/go/src/github.com/Urethramancer/name/parse.go
```

You can also combine flags, for example:
```sh
$ name -as parse.go
/Users/orb/go/src/github.com/Urethramancer/name/parse
```

