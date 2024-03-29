package request

// InputCommand for parsing input from command line
type InputCommand struct {
	// Cmd value of command, eg: add
	Cmd string
	// Value from command, eg: 10
	Value int
	// Err error helper if the input got validated
	Err error
	// Source for this command either from repeat or from user request
	Source string
}
