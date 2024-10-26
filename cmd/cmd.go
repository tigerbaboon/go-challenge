package cmd

import (
	"github.com/spf13/cobra"
)

func Commands() []*cobra.Command {
	return []*cobra.Command{
		httpCmd(),
		ex1cmd(),
		ex2Cmd(),
	}
}
