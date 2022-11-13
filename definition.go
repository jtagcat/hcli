package harg

type (
	Definitions struct {
		D       DefinitionMap
		Aliases map[string]string // map[alias slug]defSlug
	}

	DefinitionMap map[string]Definition // map[slug]; 1-character: short option, >1: long option
	Definition    struct {
		Type Type

		// For short options (1-char length), true means it's always bool
		// For long options:
		//   false: allows spaces (`--slug value` in addition to `--slug=value`)
		//   true: if "=" is not used, Type is changed to bool (or countable). Values are treated as bools, if strconv.ParseBool says so.
		// If bool is encountered after value, ErrBoolAfterValue will be returned on parsing. Any bools before value flags will be ignored.
		AlsoBool bool

		// use Definition.Methods() to get data, #TODO:
		parsed *parsedT
	}
	parsedT struct {
		originalType Type // when AlsoBool
		found        bool
		opt          option
	}
)

