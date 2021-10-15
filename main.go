package main

import (
	"fmt"

	"github.com/hgahub/dumbi-cmd/internal/version"
	"github.com/thatisuday/commando"
)

var commit string

func main() {

	// configure commando
	commando.
		SetExecutableName("dumbi").
		SetVersion(version.String(commit)).
		SetDescription("This tool lists the contents of a directory in tree-like format.\nIt can also display information about files and folders like size, permission and ownership.")

	commando.
		Register(nil).
		SetShortDescription("displays detailed information of a directory").
		SetDescription("This command displays more information about the contents of the directory like size, permission and ownership, etc.").
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			fmt.Printf("%s: %s \n\n", version.Header, version.String(commit))
			fmt.Printf("Usage: \n\tdumbi help\n\n")
		})
	commando.
		Register("validate").
		SetShortDescription("validates the configuration files").
		SetDescription("Validate runs checks that verify whether a configuration is syntactically valid.").
		AddFlag("json", "produce output in a machine-readable JSON format", commando.Bool, nil).
		SetAction(func(args map[string]commando.ArgValue, flags map[string]commando.FlagValue) {
			fmt.Printf("validate...\n\n")

			// print arguments
			for k, v := range args {
				fmt.Printf("arg -> %v: %v(%T)\n", k, v.Value, v.Value)
			}
		})

	// parse command-line arguments
	commando.Parse(nil)

}
