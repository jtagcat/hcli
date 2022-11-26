## Full Spec
- `--` ends all arguments. [^TestParseDoubledash]
- `-` by itself is a plain argument. [^TestParseNilDefs]
- Type Boolean:
    - `Count()`: Equal to the count of consecutive true values read from right/last [^TestParseCount]
    - Can never have a value (`--foo=true`, `-f false`). Set true: `--foo`, `-f`, false: `---foo`, `--f`, `-xyz-f`
    - `AlsoBool` is ignored. [^TestDefinitionNormalize]
- Prefix `--` and at least 2 UTF-8 characters means long options follow. [^TestAliasParse]
    - Long options keys are case insensitive. [^TestParseLongOptEat]
    - `=` or ` ` (space) are a delimiter in specifying values. (`--foo=bar`; `--foo=bar`, `--foo bar`) [^TestAliasParse]
      - Values parsable as options are parsed as such (`--foo --bar`: `foo != "--bar"`) [^TestParseLongOptEat]
    - Prefix `---` for Type Boolean negates it. (`---foo`) [^TestParseShortBoolOpt], [^TestParseLongOptAlsoBool], [^TestParseError]
    - `AlsoBool` treats a valueless option as a bool. [^TestParseLongOptAlsoBool]
        - Space-seperated syntax for values is unavailable. (invalid: `--foo value`) [^TestParseLongOptAlsoBool]
        - Values are always parsed as values. (`--foo=true` is string `true`, not value true) [^TestParseLongOptAlsoBool]
        - Given multiple mixed bool/value options, bools before values are ignored, and bools after value error. [^TestParseLongOptAlsoBool]
- Prefix `-` means short options follow.
    - Short options keys are 1 UTF-8 character, case sensitive. [^TestParseShortOptEat]
      - Short option keys can't start with a digit (0..9) (for ergonomics). [^TestDefinitionDigits]
    - Short options can be clustered after the prefix. (`-abc` = `-a -b -c`) [^TestParseShortBoolOpt], [^TestParseCount]
    - Preceeding `-` negates the following bool, otherwise ignored. (`--a` a:`false`; `-a-bc` a:`true` b:`false` c:`true`) [^TestParseShortBoolOpt], [^TestParseCount]
        - If `-` is used for the first short option, short options can't be clustered. (invalid:`--ab`; invalid:`--a-b` (seen as long options)) [^TestParseShortBoolOpt]
    - Non-bools take arguments until a space or from the next argument. (`-ovalue`, `-o value`) [^TestParseShortOptEat]
      - When not using space between the key and value, nothing and `=` is allowed as a delimiter (`-oval` → o:`val`, `-o=--val` → o:`--val`, `-o =val` → o:`=val`). [^TestParseShortOptEat]
      - When not using `=` as delimiter, values that could be parsed as option keys are parsed as such. (`-o -c`, o:`""`) [^TestParseShortOptEat]
    - `AlsoBool` is ignored, short options are always treated as Type. [^TestDefinitionNormalize]
- The Parser parses until any of the chokes are found. (`--foo xyz choke --bar xyz choke`: only `foo` is parsed, chokeReturn:`choke --bar xyz choke`) [^TestParseNilDefs]
    - Chokes are matched case insensitive. [^TestParseNilDefs]
    - After a choke is found, the choke and any unparsed arguments are returned on chokeReturn. [^TestParseNilDefs]
    - Chokes are not detected after arguments are ended (`--`) (no choking:`-- choke`). [^TestParseDoubledash]
    - Chokes are not detected as part of options (`--foo choke` `-o choke`) [^TestParseLongOptEat], [^TestParseShortOptEat]

[^TestParseNilDefs]: Tested by `TestParseNilDefs()`
[^TestParseLongOptEat]: Tested by `TestParseLongOptEat()`
[^TestParseShortOptEat]: Tested by `TestParseShortOptEat()`
[^TestParseDoubledash]: Tested by `TestParseDoubledash()`
[^TestParseLongOptAlsoBool]: Tested by `TestParseLongOptAlsoBool()`
[^TestParseShortBoolOpt]: Tested by `TestParseShortBoolOpt()`
[^TestDefinitionNormalize]: Tested by `TestDefinitionNormalize()`
[^TestParseCount]: Tested by `TestParseCount()`
[^TestParseError]: Tested by `TestParseError()`
[^TestDefinitionDigits]: Tested by `TestDefinitionDigits()`

### Additions compared to GNU:
Based on https://www.gnu.org/software/libc/manual/html_node/Argument-Syntax.html, the following has been added:

- def.Sl(): (`--foo bar --foo baz` foo:{`bar`,`baz`})
- Space seperator (lookahead) in long options.
    - `AlsoBool`: Disallows space seperator, allows mixed bool (`--foo`) and valueful (`--foo=value`) definitions.
- Negating short options: adding `-` before a short option means `false` (`--f`, `-b-f`).
- Negating long options: adding `-` before a long option means `false` (`---foo`).
- Chokes parse until a keyword is found. This allows crafting subcommands, and global-local options.
