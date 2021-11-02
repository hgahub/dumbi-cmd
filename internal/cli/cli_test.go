package cli

import (
	"github.com/hgahub/dumbi-cmd/internal"
	"github.com/hgahub/dumbi-cmd/internal/cli/command"
	"os"
	"testing"
)

func TestGeneralFormalVerification(t *testing.T) {
	tests := map[string]struct {
		input string
	}{
		"empty parameter": {input: ""},
		"one parameter":   {input: "one"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			os.Args = []string{tc.input}
			c, err := CommandLineRead(nil)

			if c != nil || err.Error() != InvalidArgumentsMsg {
				t.Fatalf("General formal error")
			}
		})
	}
}

func TestInvalidGlobalFlagOrCommand(t *testing.T) {
	os.Args = []string{"app", "-invalid"}
	c, err := CommandLineRead(nil)

	if c != nil || err.Error() != InvalidGlobalFlagOrCommandMsg {
		t.Fatalf("Global flag error")
	}
}

func TestValidGlobalFlag(t *testing.T) {
	tests := map[string]struct {
		input string
	}{
		"help":    {input: "-help"},
		"version": {input: "-version"},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			os.Args = []string{"app", tc.input}
			c, err := CommandLineRead(internal.NewVersion("", "1", ""))

			if err != nil && c.Name != command.GlobalFlagCommandName {
				t.Fatalf("Global flag error")
			}

			if c.Flags[tc.input].Name != tc.input[1:] {
				t.Fatalf("Flag error: %s", c.Flags[tc.input].Name)
			}
		})
	}
}

func TestValidCommand(t *testing.T) {
	tests := map[string]struct {
		cmd  string
		name string
	}{
		command.ValidateCommandName: {cmd: command.ValidateCommandName, name: command.ValidateCommandName},
		"fmt":                       {cmd: "fmt", name: "format"},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			os.Args = []string{"app", tc.cmd}
			c, err := CommandLineRead(nil)

			if err != nil || c == nil {
				t.Fatalf("Command error")
			}

			if c.Name != tc.name {
				t.Fatalf("Command error: %s", c.Name)
			}
		})
	}
}
