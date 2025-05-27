/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all tasks",
	Long: `Display all tasks with their stauts`,
	Run: func(cmd *cobra.Command, args []string) {
		tasks := loadTasks()
		if len(tasks) == 0 {
			fmt.Println("No Tasks found")
		}
		for i , v := range tasks {

			status := "Pending"
			if v.Done {
				status = "Done"
			}
			fmt.Printf("%v. %v [%v] \n", i+1, v.Description , status)
		}
		
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
