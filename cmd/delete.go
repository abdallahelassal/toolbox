/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete [Task Number]",
	Short: "delet a task",
	Long: `Delete a task by specifying its task number.`,
	Args: cobra.ExactArgs(1), 
	Run: func(cmd *cobra.Command, args []string) {
		index , err := strconv.Atoi(args[0])
		if err != nil || index < 1 {
			fmt.Println("invalid task number")
			return
		}
		tasks := loadTasks()
		if index > len(tasks) {
			fmt.Println("Task number out of range.")
			return
		}
		tasks = append(tasks[:index-1],tasks[index:]... )
		saveTasks(tasks)
		fmt.Printf("Task %d deleted.\n", index)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
