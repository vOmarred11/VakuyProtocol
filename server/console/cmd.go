package console

import (
	"fmt"
	"github.com/vOmarred11/VakuyProtocol/command"
)

// Command is the command that you execute on the console
type Command struct {
	// ConsoleLine is the command line that appears when you execute the command on console
	ConsoleLine string
	// ConsoleLine is the message that appears in the console when you execute the command on the console
	ConsoleOutput string
	// Commands is a slice of all available commands
	Commands []command.Command
	// Arguments is a slice of all arguments of every command
	Arguments []command.Arguments
	// Output is the action that will get taken
	Output command.Output
}

func (c *Command) CommandConsoleLine() string {
	return c.ConsoleLine
}
func (c *Command) CommandConsoleOutput() string {
	return c.ConsoleOutput
}
func (c *Command) CommandCommands() []command.Command {
	return c.Commands
}
func (c *Command) CommandArguments() []command.Arguments {
	return c.Arguments
}
func (c *Command) CommandOutput() command.Output {
	fmt.Sprintf("%s | %s", c.ConsoleLine, c.ConsoleOutput)
	return c.Output
}
