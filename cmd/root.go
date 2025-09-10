package cmd

import (
	"github.com/spf13/cobra"
)

// The command to run the main program (Our program is called "task").
var RootCmd = &cobra.Command{
	Use:   "task",
	Short: "Task is a CLI Task Manager.",
	// Run: func(cmd *cobra.Command, args []string) { },
}

// init to be called before main runs; to use it later.
func init() {
	RootCmd.CompletionOptions.DisableDefaultCmd = true
}
