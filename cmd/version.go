package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Version  string
	Revision string
)

func newVersionCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print version",
		Run:   runVersionCmd,
	}
	return cmd
}

func runVersionCmd(_ *cobra.Command, _ []string) {
	fmt.Printf("geso version: %s, revision: %s\n", Version, Revision)
}
