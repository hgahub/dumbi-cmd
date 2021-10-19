package cli

// data type declarations
const (
	Bool = iota
	Int
	String
)

type Flag struct {
	// description of the flag
	description string

	// data type of the flag value
	dataType int
}

type Command struct {
	// description of the command
	description string

	// flags to parse from the command-line arguments
	flags map[string]*Flag
}
