package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResolveCliCommand(t *testing.T) {
	var args []string
	var cmd command
	asserts := assert.New(t)

	args = []string{"./cli"}
	cmd = resolveCliCommand(args)
	asserts.Equal(command{
		name: CmdHelp,
	}, cmd, "Call without arguments must return result of help handler")

	args = []string{"./cli", "make:user"}
	cmd = resolveCliCommand(args)
	asserts.Equal(command{
		name:  CmdMakeUser,
		flags: map[string]string{},
	}, cmd, "Call with known argument must return proper command")

	args = []string{"./cli", "make:unknown", "--name=John", "./file.txt"}
	cmd = resolveCliCommand(args)
	asserts.Equal(command{
		name:     "make:unknown",
		flags:    map[string]string{"name": "John"},
		argument: "./file.txt",
	}, cmd, "Command must contain name, flags and argument that passed to cli")
}
