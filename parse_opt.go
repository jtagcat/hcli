package harg

import "fmt"

func (defM *DefinitionMap) toEmptyOptM() OptionsMap {
	return OptionsMap{} // TODO:
}

type (
	OptionsMap map[string]Option
	Option     struct {
		// TODO: ???
	}
)

// TODO:
// type (
// 	ArgBoolean struct {
// 		Defined bool // if found in args
// 		Value   bool // result
// 	}
// )
// type (
// 	RawArgs map[string]Arg[string] // string: identifier slug

//	Arg[T any] struct {
//		Defined bool
//		HasData bool
//		Data    []T
//	}

func (optM *OptionsMap) parseOptionContent(def *Definition, value *string, valueFound *bool) error {
	return fmt.Errorf("not implemented") // TODO:
}
