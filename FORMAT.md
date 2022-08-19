## Full Spec:
- `--` ends all arguments.
- `-` is a plain argument.
- Where multiple arguments conflict, the following happens:
    - default: Everything but the last is ignored.
    - def.Count: they are counted
    - def.Slice: result is a slice of values (takes priority over def.Slice)
        - AlsoBoolean: Input can't mix boolean and valueful in the same parse operation.
- Prefix `--` and at least 2 utf8 characters means long options follow.
    - Long options are case insensitive.
    - `=` is a single delimiter in specifying values. (`--foo=bar`; `--foo=false`)
    - A space may be used to specify values for non-boolean types. (`--foo bar`)
    - AlsoBoolean treats a valueless valueful option as a boolean. (`--foo`; `--foo=value`)
        - Space-seperated syntax is unavailable. (invalid:`--foo value`)
        - Booleans in values are not parsed. (`--foo=true` is string value `"true"`)
- Prefix `-` means short options follow.
    - Short options are 1 utf8 character, case sensitive.
    - Short options can be clustered after the prefix. (`-abc` a:`true` b:`true` c:`true`)
    - Preceeding `-` negates the following boolean, otherwise ignored. (`--a` a:`false`; `-a-bc` a:`true` b:`false` c:`true`)
        - If `-` is used for the first short option, short options can't be clustered. (invalid:`--ab`; invalid:`--a-b` (seen as long options))
    - Non-booleans take arguments until space or from the next argument. (`-aovalue`, `-ao value` a:`false` o:`value`)
    - Definitions: AlsoBoolean: Short option is always treated as a boolean.
- The Parser only parses everything left of the first choke found. (`--foo xyz choke --bar xyz choke` foo:`xyz`, chokeReturn:`choke --bar xyz choke`)
    - Chokes are defined in an input to the parser.
    - Chokes are matched case insensitive.
    - The choke and any following arguments are returned on chokeReturn.
    - Chokes are not detected after `--` (no choking:`-- choke`).
    - Chokes are not detected as part of options (no choking:`--foo choke` foo:`choke`; no choking:`-o choke -b` o:`choke` b:`true`)

### Additions:
Based on https://www.gnu.org/software/libc/manual/html_node/Argument-Syntax.html, the following has been added:

- def.Slice: (`--foo bar --foo baz` foo:{`bar`,`baz`})
- Space seperator in long options.
    - AlsoBoolean: Disallows space seperator, allows mixed boolean (`--foo`) and valueful (`--foo=value`) definitions.
- Negative short options: adding `-` before a short option means `false`.
- Chokes enables parsing until a keyword is found. This allows crafting global-local-superglobal-whatever options.
