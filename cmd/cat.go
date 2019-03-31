package cmd

import (
	"github.com/spf13/cobra"
)

func newCatCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:           "cat",
		Short:         "Get usage cat command",
		SilenceErrors: true,
		SilenceUsage:  false,
	}

	cmd.AddCommand(
		newCatIndicesCmd(),
		newCatAliasesCmd(),
	)
	return cmd
}
