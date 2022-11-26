# harg

GNU-compatible feature-complete Go argument parsing. See [FORMAT.md](FORMAT.md) for full specification.

See also: https://github.com/urfave/cli/issues/833#issuecomment-1312834335

Name is a play on https://git.meatballhat.com/x/argh; may also stand for 'human-friendly arguments' and 'harrrggghh!' 🏴‍☠️, finally something good for Go arguments!

## Left to do:
- Review codebase errors, flow clarity, and overall; `golangci-all`
- Example use for pkg.go.dev / README.md.
- `strconv` says `t` and `f` are bools, and `y`, `n`, `yes`, `no` aren't. Are they?

### Next up:
- henv: Environment variables
- [`urfave/cli@v3`](https://github.com/urfave/cli)?

### Niceties:
- Definition-based shell completions
- `hyaml`: `yaml`?
- ~~Code generation?~~
