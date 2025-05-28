/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"
	"github.com/spf13/cobra"
)

// doneCmd represents the done command
var doneCmd = &cobra.Command{
	Use:   "done [Tasks Number]",
	Short: "Mark a task as done ",
	Long: `Mark a task as done by specifying its task number.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		index , err := strconv.Atoi(args[0])
		if err != nil || index < 1{
			fmt.Println("Invalid task number")
			return 
		} 
		tasks := loadTasks()
		if index > len(tasks) {
			fmt.Println("Tasks number of range ")
			return
		}
		tasks[index-1].Done = true
		saveTasks(tasks)
		fmt.Printf("Tasks %v is Done \n", index )
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// doneCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// doneCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
