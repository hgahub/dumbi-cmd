package command

// data type declarations
const (
	Bool = iota
	Int
	String
)

const GlobalFlagCommandName = "global"

type Flag struct {
	// Name of the flag
	Name string

	// Description of the flag
	Description string

	// data type of the flag value
	DataType int
}

type Command struct {
	// Name of the command
	Name string

	// Description of the command
	Description string

	// Flags to parse from the command-line arguments
	Flags map[string]*Flag
}

type CLICommand interface {
	run() int
}

func (c *Command) Execute() int {
	var cliCommand CLICommand
	meta := NewMeta()

	switch c.Name {
	case ValidateCommandName:
		cliCommand = &ValidateCommand{meta: meta}
	default:
		return 126
	}

	return cliCommand.run()
}
