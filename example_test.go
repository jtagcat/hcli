package harg_test

import (
	"fmt"
	"log"
	"strings"

	"github.com/jtagcat/harg"
)

func ExampleParse() {
	kOne, kTwo, kThree := "o", "t", "three"
	defs := harg.Definitions{
		kOne:   {Type: harg.String},
		kTwo:   {Type: harg.Bool},
		kThree: {Type: harg.Duration},
	}

	osArgs := strings.Split("programName hello -to foo -o bar --three 5s --t -t -t world", " ")

	args, _, err := defs.Parse(osArgs[1:], nil)
	if err != nil {
		log.Fatalf("parsing command-line arguments: %e", err)
	}

	fmt.Println(args) // [hello world]

	sl, ok := defs[kOne].SlString() // ok: is valid and set
	if ok {
		fmt.Println(sl) // [foo bar]
	}
	s, ok := defs[kOne].String()
	if ok {
		fmt.Println(s) // bar
	}

	two, _ := defs[kTwo].SlBool()
	fmt.Println(two) // [true false true true]

	count, _ := defs[kTwo].Count() // how many true in a row
	fmt.Println(count)             // 2

	dur, _ := defs[kThree].Duration()
	fmt.Println(dur) // 5s
}
