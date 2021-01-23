package cmd

import (
	"github.com/MadhavJivrajani/halp/morse"
	"github.com/spf13/cobra"
)


// resetCmd represents the reset command
var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset the capslock LED to an OFF state",
	Long: `
Doen't matter what arguments you give it, it will reset.

Syntax:
halp reset`,

	RunE: func(cmd *cobra.Command, args []string) error {
		path, _ := cmd.Flags().GetString("path")
		return morse.UpdateState(path, 0)
	},
}

func init() {
	rootCmd.AddCommand(resetCmd)
}
