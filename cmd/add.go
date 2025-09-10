package cmd

import (
	"fmt"
	"os"
	"strings"
	"task/db"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to your task list.",
	Run: func(cmd *cobra.Command, args []string) {
		// in case task has more than one word -> join into sentence.
		task := strings.Join(args, " ")

		// create the task in db and handle errors.
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("Something went wrong:", err)
			os.Exit(1)
		}
		fmt.Printf("Added \"%s\" to your task list.\n", task)
	},
}

// init to be called before main runs; to use it later.
func init() {
	RootCmd.AddCommand(addCmd)
}
