package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:           "geso",
	Short:         "Elasticsearch CLI",
	SilenceErrors: true,
	SilenceUsage:  true,
}

func init() {
	rootCmd.PersistentFlags().StringP("url", "", "http://localhost:9200", "Elasticsearch endpoint URL")

	prepareSubCommands()
}

// add sub commands
func prepareSubCommands() {
	rootCmd.AddCommand(
		newVersionCmd(),
		newCatCmd(),
	)
}

func Exec() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Printf("%+v\n", err)
		os.Exit(1)
	}
}
