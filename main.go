package main

import (
	"app/cmd"

	"github.com/spf13/cobra"
)

func main() {
	cmds := &cobra.Command{
		Use:   "cmd",
		Short: "List all commands",
	}

	cmds.AddCommand(cmd.Commands()...)

	cmds.Execute()
}
