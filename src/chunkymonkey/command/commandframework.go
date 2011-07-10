package command

import (
	"strings"
	"os"
)

var ErrCmdExists = os.NewError("The command already exists.")

// The CommandFramework handles all message based commands.
// It uses channels to safly handle multiple calls.
type CommandFramework struct {
	prefix string // The command prefix befor every command.
	cmds   map[string]*Command
}

// Creates a new CommandFramework and starts the update process.
func NewCommandFramework(prefix string) *CommandFramework {
	cf := &CommandFramework{prefix: prefix}
	cmds := getCommands()
	commandHelp := NewCommand(helpCmd, helpDesc, helpUsage, func(msg string, cmdHandler ICommandHandler) {
		cmdHelp(msg, cf, cmdHandler)
	})
	cmds[helpCmd] = commandHelp
	cmds[helpShortCmd] = commandHelp
	cf.cmds = cmds
	return cf
}

func (cf *CommandFramework) Prefix() string {
	return cf.prefix
}

func (cf *CommandFramework) Commands() map[string]*Command {
	return cf.cmds
}

func (cf *CommandFramework) Process(message string, cmdHandler ICommandHandler) {
	if len(message) < 2 || message[0:len(cf.prefix)] != cf.prefix {
		return
	}
	attr := strings.Split(message, " ", -1)
	trigger := attr[0][1:]
	if cmd, ok := cf.cmds[trigger]; ok {
		cmd.Callback(message, cmdHandler)
	}
}