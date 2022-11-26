## Full Spec
```
(os.Args) arguments = shortOptions + longOptions + parsedArgs + chokeReturn
```

- `--` ends all arguments. [^TestParseDoubledash]
- `-` is a plain argument. [^TestParseNilDefs]
- Type Boolean:
    - `Count()`: Equal to the count of consecutive true values read from right/last [^TestParseCount]
    - `AlsoBool` is ignored. [^TestDefinitionNormalize]
- Prefix `--` and at least 2 UTF-8 characters means long options follow. [^TestAliasParse]
    - Long options are case insensitive. [^TestParseLongOptEat]
    - `=` is a single delimiter in specifying values. (`--foo=bar`; `--foo=false`) [^TestAliasParse]
    - A space may be used to specify values for non-bool types. (`--foo bar`) [^TestAliasParse]
      - Values that could be parsed as option keys (`--foo --baz`) are parsed as keys (2 empty), not as values (foo:`--bar`). [^TestParseLongOptEat]
    - `AlsoBool` treats a valueless valueful option as a bool. (`--foo`; `--foo=value`) [^TestParseLongOptAlsoBool]
        - Space-seperated syntax is unavailable. (invalid:`--foo value`) [^TestParseLongOptAlsoBool]
        - Bools in values are parsed as booleans. (`--foo=true` is bool, not string "true") [^TestParseLongOptAlsoBool]
        - Given multiple mixed bool/value options, bools before values are ignored, and bools after value error. [^TestParseLongOptAlsoBool]
- Prefix `-` means short options follow.
    - Short options are 1 utf8 character, case sensitive. [^TestParseShortOptEat]
    - Short options can be clustered after the prefix. (`-abc` a:`true` b:`true` c:`true`) [^TestParseShortBoolOpt], [^TestParseCount]
    - Preceeding `-` negates the following bool, otherwise ignored. (`--a` a:`false`; `-a-bc` a:`true` b:`false` c:`true`) [^TestParseShortBoolOpt], [^TestParseCount]
        - If `-` is used for the first short option, short options can't be clustered. (invalid:`--ab`; invalid:`--a-b` (seen as long options)) [^TestParseShortBoolOpt]
    - Non-bools take arguments until space or from the next argument. (`-aovalue`, `-ao value` a:`false` o:`value`) [^TestParseShortOptEat]
      - When not using space between value, nothing and `=` is allowed as a delimiter (`-oval` → o:`val`, `-o=--val` → o:`--val`, `-o =val` → o:`=val`). [^TestParseShortOptEat]
      - Values that could be parsed as option keys (`-ao -c`) are parsed as keys (o:empty), not as values (`-o=--bar` → o:`--bar`). [^TestParseShortOptEat]
    - `AlsoBool` is ignored, option is always treated as Type. [^TestDefinitionNormalize]
- The Parser only parses everything left of the first choke found. (`--foo xyz choke --bar xyz choke` foo:`xyz`, chokeReturn:`choke --bar xyz choke`) [^TestParseNilDefs]
    - Chokes are matched case insensitive. [^TestParseNilDefs]
    - The choke and any following arguments are returned on chokeReturn. [^TestParseNilDefs]
    - Chokes are not detected after `--` (no choking:`-- choke`). [^TestParseDoubledash]
    - Chokes are not detected as part of options (no choking:`--foo choke` foo:`choke`; no choking:`-o choke -b` o:`choke` b:`true`) [^TestParseLongOptEat],[^TestParseShortOptEat]

[^TestParseNilDefs]: Tested by `TestParseNilDefs()`
[^TestParseLongOptEat]: Tested by `TestParseLongOptEat()`
[^TestParseShortOptEat]: Tested by `TestParseShortOptEat()`
[^TestParseDoubledash]: Tested by `TestParseDoubledash()`
[^TestParseLongOptAlsoBool]: Tested by `TestParseLongOptAlsoBool()`
[^TestParseShortBoolOpt]: Tested by `TestParseShortBoolOpt()`
[^TestDefinitionNormalize]: Tested by `TestDefinitionNormalize()`
[^TestParseCount]: Tested by `TestParseCount()`

### Additions compared to GNU:
Based on https://www.gnu.org/software/libc/manual/html_node/Argument-Syntax.html, the following has been added:

- def.Sl(): (`--foo bar --foo baz` foo:{`bar`,`baz`})
- Space seperator (lookahead) in long options.
    - `AlsoBool`: Disallows space seperator, allows mixed bool (`--foo`) and valueful (`--foo=value`) definitions.
- Negative short options: adding `-` before a short option means `false`.
- Chokes enables parsing until a keyword is found. This allows crafting global-local-superglobal-whatever options.
