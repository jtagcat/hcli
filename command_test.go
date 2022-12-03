package hcli_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/jtagcat/hcli"
)

func TestRun(t *testing.T) {
	app := hcli.Command{
		Flags: hcli.Flags{
			"foo": hcli.StringFlag{
				Options: []string{"foo", "f"}, Env: []string{"FOOBEANS"},
				Source: hcli.OptEnv, Default: "brr", Condition: hcli.Defined,
				Usage: "its for foo energy bars",
			},
		},

		Action: func(ctx hcli.Context) (_ error, exitCode int) {
			s, ok := ctx.SlString("foo")

			return fmt.Errorf("not implemented; foo was set: %b, with value: %q", ok, s), 1
		},
	}

	exitCode := app.Run(os.Args[0], os.Args[1:])
	os.Exit(exitCode)
}
