package hcli_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/jtagcat/hcli"
	"github.com/jtagcat/hcli/harg"
)

func TestRun(t *testing.T) {
	app := hcli.Command{
		Flags: []hcli.Flag{
			hcli.StringFlag{
				Options: []string{"foo", "f", "bar"}, Env: []string{"FOOBEANS"},
				Source: hcli.OptEnv, Default: "brr", Condition: hcli.Defined,
				Usage: "its for foo energy bars",
			},
			hcli.StringFlag{Env: []string{"MEOW"}, Condition: hcli.NotDefault},
			hcli.BoolFlag{Env: []string{"ACKNOWLEDGE_RISKS"}, Condition: func(_ any, def *harg.Definition) error {
				b, _ := def.Bool()
				if b {
					return nil
				}

				return fmt.Errorf("program will not run unless ACKNOWLEDGE_RISKS is set to true") // --help will be called
			}},
		},

		Action: func(ctx hcli.Context) (_ error, exitCode int) {
			s, ok := ctx.SlString("bar")
			fmt.Println("--foo, -f, or FOOBEANS was set: %b, with value: %q", ok, s)

			b, ok := ctx.Bool("MEOW")
			fmt.Println("MEOW was set: %b, with value: %b", ok, b)

			return fmt.Errorf("not implemented"), 1
		},
	}

	exitCode := app.Run(os.Args[0], os.Args[1:])
	os.Exit(exitCode)
}
