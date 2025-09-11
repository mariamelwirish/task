package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mariamelwirish/task/db"

	"github.com/spf13/cobra"
)

var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit your task title.",
	Run: func(cmd *cobra.Command, args []string) {
		// parse id
		id, err := strconv.Atoi(args[0])
		tasks, _ := db.AllTasks()
		if err != nil || id <= 0 || id > len(tasks) {
			fmt.Printf("Invalid task number: %s. Error:\n", args[0], err)
			return
		}

		// join remaining args as the new text
		newTask := strings.TrimSpace(strings.Join(args[1:], " "))
		if newTask == "" {
			fmt.Printf("New task cannot be empty!")
			return
		}

		// map the user id to DB key and delete.
		task := tasks[id -1]
		err2 := db.UpdateTask(task.Key, newTask)
		if err2 != nil {
			fmt.Printf("Failed to edit \"%d\". Error: %s\n", id, err2)
		} else {
			fmt.Printf("Changed task \"%d\" to %s.\n", id, newTask)
		}
		
	},
}

// init to be called before main runs; to use it later.
func init() {
	RootCmd.AddCommand(editCmd)
}
