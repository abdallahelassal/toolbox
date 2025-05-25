/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"github.com/spf13/cobra"
)

type Task struct {
	Description string `json:"description"`
	Done bool `json:"done"`
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [ task decription ]",
	Short: "add a new task",
	Long: `Add a new task with a description to your task list.`,
	Args: cobra.ExactArgs(1), // Ensure exactly one argument is provided
	Run: func(cmd *cobra.Command, args []string) {
		description  := args[0]
		tasks := loadTasks()
		tasks = append(tasks, Task{Description: description, Done: false})
		saveTasks(tasks)
		fmt.Printf("Tasks added: %s\n", description)
	},
}

	func loadTasks() []Task{
		file,err := os.Open("tasks.json")
		if os.IsNotExist(err){
			return []Task{}
		}else if err != nil {
			fmt.Println("Error loading tasks:", err)
			return []Task{}
		}
		defer file.Close()
		var tasks []Task
		json.NewDecoder(file).Decode(&tasks)
		return tasks
	}
	func saveTasks(tasks []Task){
		file,err := os.Create("tasks.json")
		if err !=nil {
			fmt.Println("Error saving tasks:", err)
			
		}
		defer file.Close()
		json.NewEncoder(file).Encode(tasks)
		fmt.Println("Tasks saved successfully.")
	}



// func init() {
// 	rootCmd.AddCommand(addCmd)

// 	// Here you will define your flags and configuration settings.

// 	// Cobra supports Persistent Flags which will work for this command
// 	// and all subcommands, e.g.:
// 	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

// 	// Cobra supports local flags which will only run when this command
// 	// is called directly, e.g.:
// 	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
// }
