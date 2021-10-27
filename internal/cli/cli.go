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

const InvalidArgumentsMsg = "invalid arguments"
const InvalidGlobalFlagOrCommandMsg = "invalid global flag or command"

// Collection of usable commands
var commands map[string]Command

// Initialize the commands map
func init() {
	commands = map[string]Command{
		"validate": {Name: "validate", Description: "validates the configuration files"},
		"fmt":      {Name: "format", Description: "rewrite Dumbi configuration files to a canonical format and style"},
	}
}

func CommandLineRead(version *internal.Version) (*Command, error) {
	// General formal verification
	if len(os.Args) < 2 {
		fmt.Printf(getHelpText())
		return nil, fmt.Errorf(InvalidArgumentsMsg)
	}

	cmd, isCmd := commands[os.Args[1]]

	if !isCmd {
		if f, err := globalFlagVerification(version); err == nil {
			return &Command{Name: "global", Flags: map[string]*Flag{os.Args[1]: f}}, nil
		} else {
			msg := fmt.Sprintf("Dumbi has no command named \"%s\".\n\n", os.Args[1])
			msg += fmt.Sprintln("To see all of Dumbis's commands, run:")
			msg += fmt.Sprintln("  dumbi -help")

			fmt.Printf(msg)
			return nil, fmt.Errorf(InvalidGlobalFlagOrCommandMsg)
		}
	} else {
		return &cmd, nil
	}
}

func globalFlagVerification(version *internal.Version) (*Flag, error) {
	var versionString string

	if version != nil {
		versionString = version.String()
	}

	switch os.Args[1] {
	case "-help":
		fmt.Printf(getHelpText())
		return &Flag{Name: "help", DataType: String}, nil
	case "-version":
		fmt.Printf("Dumbi %s\n", versionString)
		return &Flag{Name: "version", DataType: String}, nil
	default:
		return nil, fmt.Errorf("unknown global flag: %s", os.Args[1])
	}
}

func getHelpText() string {
	commandsHelp := ""
	for k, cmd := range commands {
		commandsHelp += fmt.Sprintf("  -%-12s %s\n", k, cmd.Description)
	}

	return fmt.Sprintf(helpText, commandsHelp)
}
