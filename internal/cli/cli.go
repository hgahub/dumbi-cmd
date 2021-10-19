package cli

import (
	"fmt"
	"github.com/hgahub/dumbi-cmd/internal"
	"os"
)

const helpText = `
Usage: dumbi [options] <command> [args]

The available commands for execution are listed below.

Commands:
%s

Options (use these before the subcommand, if any):
  -help         Show this help output, or the help for a specified subcommand.
  -version      Displays the version of Dumbi.
`

var commands map[string]Command
var versionString string

func init() {
	commands = map[string]Command{
		"validate": {description: "validates the configuration files"},
		"fmt":      {description: "rewrite Dumbi configuration files to a canonical format and style"},
	}
}

func Parse(version *internal.Version) (*Command, error) {
	versionString = version.String()
	if len(os.Args) < 2 {
		fmt.Printf(helpText, "dd")
		return nil, nil
	}

	cmd, isCmd := commands[os.Args[1]]

	if !isCmd {
		if checkGlobalFlag() {
			return nil, nil
		}
		msg := fmt.Sprintf("Dumbi has no command named \"%s\".\n\n", os.Args[1])
		msg += fmt.Sprintln("To see all of Dumbis's commands, run:")
		msg += fmt.Sprintln("  dumbi -help")

		return nil, fmt.Errorf(msg)
	}

	return &cmd, nil
}

func checkGlobalFlag() bool {
	switch os.Args[1] {
	case "-help":
		fmt.Printf(helpText, "dd")
	case "-version":
		fmt.Printf("Dumbi %s\n", versionString)
	default:
		return false
	}

	return true
}
