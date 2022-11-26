# harg

[![Go Reference](https://pkg.go.dev/badge/github.com/jtagcat/harg.svg)](https://pkg.go.dev/github.com/jtagcat/harg)

GNU-compatible feature-complete Go argument parsing. See [FORMAT.md](FORMAT.md) for full specification.

See also: https://github.com/urfave/cli/issues/833#issuecomment-1312834335

Name is a play on https://git.meatballhat.com/x/argh; may also stand for 'human-friendly arguments' and 'harrrggghh!' 🏴‍☠️, finally something good for Go arguments!

### Next up:
- henv: Environment variables
- [`urfave/cli@v3`](https://github.com/urfave/cli)?

### Niceties:
- Definition-based shell completions
- `hyaml`: `yaml`?
- ~~Code generation?~~

## Code flow
1. [`definition.go`](definition.go): definition structs
1. [`parse.go`](parse.go): main routine, splits to short/long option
1. [`parse_option.go`](parse_option.go): short and long option parsing
1. [`option_parse.go`](option_parse.go): parsing values to definitions
1. [`option_set.go`](option_set.go): typed structs
1. [`option_get.go`](option_get.go): typed structs, public functions for retrieving values.
