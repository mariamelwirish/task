package cmd

import (
	"fmt"
	"os"
	"github.com/mariamelwirish/task/db"

	"github.com/spf13/cobra"
)

// list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all of your tasks.",
	Run: func(cmd *cobra.Command, args []string) {
		// Call AllTasks() function implemented earlier.
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong:", err)
			os.Exit(1)
		}

		// Check if there are no tasks.
		if(len(tasks) == 0) {
			fmt.Println("You have no tasks to complete! :D")
			return
		}

		// List tasks.
		fmt.Println("You have the following tasks: ")
		for i, task := range tasks {
			fmt.Printf("%d. %s\n", i+1, task.Value)
		}
		
	},
}

// init to be called before main runs; to use it later.
func init() {
	RootCmd.AddCommand(listCmd)
}
