package hcli_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/jtagcat/hcli"
	"github.com/jtagcat/hcli/harg"
	"log/slog"
)

func TestRun(t *testing.T) {
	app := hcli.Command{
		Flags: []hcli.Flag{
			// define a global (within a sub-command tree)
			&hcli.BoolFlag{
				Level: hcli.Global, Options: []string{"acknowledge-risks"},
				Condition: func(_ []bool, def *harg.Definition) error {
					b, _ := def.Bool()
					if b {
						return nil
					}

					return fmt.Errorf("program will not run unless --acknowledge-risks is set (to true)") // --help will be called
				},
			},
			// opt-in implementation / usage in children:
			&hcli.ChildFlag{Parent: "acknowledge-risks"}, // for Child: opt or env doesn't matter, everything works

			// locals (default Level)
			&hcli.StringFlag{
				Options: []string{"foo", "f", "bar"}, Env: "FOOBEANS",
				Default: []string{"brr"}, Condition: hcli.Defined[string],
				Usage: "its for foo energy bars",
			},
			&hcli.StringFlag{Env: "MEOW", Condition: hcli.NotDefault[string]},
		},

		Action: func(ctx context.Context, args []string, opts harg.Definitions, log *slog.Logger) (exitCode int) {
			sl, ok := opts["bar"].SlString()
			log.Debug("energy bars", slog.Any("value", sl), slog.Bool("ok", ok),
				slog.String("name", "--foo"), slog.String("name", "-f"), slog.String("name", "FOOBEANS"))

			b, ok := opts["MEOW"].Bool()
			log.Info("MEOW", nil, slog.Bool("value", b), slog.Bool("ok", ok))

			log.Error("not implemented", nil)
			return 1
		},
	}

	exitCode := app.Run(context.Background(), os.Args[0], os.Args[1:])
	os.Exit(exitCode)
}
