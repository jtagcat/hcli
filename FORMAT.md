## Full Spec:
- `--` ends all arguments.
- `-` is a plain argument.
- Where multiple arguments conflict, the following happens:
    - Default: Everything but the last is ignored.
    - def.Count: they are counted
    - def.Slice: result is a slice of values
- Prefix `--` and at least 2 utf8 characters means long options follow.
    - Long options are case insensitive.
    - To specify values, `=` is the delimiter. `--foo=bar`, `--foo=false`
    - To specify values for non-boolean Types, space may be used. `--foo bar`
    - Definitions: AlsoBoolean: Space-seperated syntax is unavailable, booleans are not parsed, `--foo=true` is string value `"true"` (`--foo`, `--foo=value`)
- Prefix `-` means short options follow.
    - Short options are 1 utf8 character, case sensitive.
    - Short options can be clustered after the prefix. `-abc` (a:`true` b:`true` c:`true`)
    - Preceeding `-` negates the following boolean, otherwise ignored. `--a` (a:`false`), `-a-bc` (a:`true`, b:`false`, c:`true`)
        - If `-` is used for the first short option, short options can't be clustered. (invalid: `--ab`, `--a-b` (seen as long options))
    - Non-booleans take arguments until space or as the next argument. `-aovalue`, `-ao value` (`false`, `value`)
    - Definitions: AlsoBoolean: Short option is always treated as a boolean.
- Chokes: #TODO:

### Additions:
Based on https://www.gnu.org/software/libc/manual/html_node/Argument-Syntax.html

- Slice definition: (`--include a --include b`) instead of keeping the last (include:`b`), specifying values multiple times results in a slice of values (include:{`a`,`b`}).
- Space seperator for long arguments.
    - AlsoBoolean: Disallows space seperator, allows mixed boolean (`--foo`) and valueful (`--foo=value`) definitions.
- Negative short options. Adding `-` before a short option means false.
- Chokes: allowing partial parsing until keyword (see full spec)
