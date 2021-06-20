/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strconv"
	"strings"
	"toDoList/todolist"
)

// todolistRootCmd represents the todolist command
var todolistRootCmd = &cobra.Command{
	Use:   "todolist",
	Short: "A To Do List CLI",
	Long:  "To Do List CLI is a tool to organise your daily tasks with add, remove, done and undone functionalities",
}

// todolistAddCmd represents the todolist command
var todolistAddCmd = &cobra.Command{
	Use:                   "add <task_name>",
	Short:                 "Add task to the list",
	Long:                  "Add task to the list",
	Args:                  cobra.MaximumNArgs(25),
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		if strings.TrimSpace(task) != "" {
			todolist.Add(task)
		} else {
			fmt.Println("Task cannot be empty. Try again")
		}
	},
}

// todolistCleanupCmd represents the todolist command
var todolistCleanupCmd = &cobra.Command{
	Use:                   "cleanup",
	Short:                 "Cleanup done tasks",
	Long:                  "Cleanup done tasks",
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		todolist.Cleanup()
	},
}

// todolistDoneCmd represents the todolist command
var todolistDoneCmd = &cobra.Command{
	Use:                   "done <task_id>",
	Short:                 "Mark task as done",
	Long:                  "Mark task as done",
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		taskId, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("%v not a valid integer. Try again\n", args[0])
			os.Exit(1)
		}
		todolist.Done(taskId - 1) // -1 because slice index starts from 0
	},
}

// todolistListCmd represents the todolist command
var todolistListCmd = &cobra.Command{
	Use:                   "list",
	Short:                 "List all tasks still to do",
	Long:                  "List all tasks still to do",
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		todolist.List()
	},
}

// todolistUndoneCmd represents the todolist command
var todolistUndoneCmd = &cobra.Command{
	Use:                   "undone <task_id>",
	Short:                 "Mark task as not done",
	Long:                  "Mark task as not done",
	Args:                  cobra.ExactArgs(1),
	DisableFlagsInUseLine: true,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		taskId, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("%v not a valid integer. Try again\n", args[0])
			os.Exit(1)
		}
		todolist.UnDone(taskId - 1) // -1 because slice index starts from 0
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(todolistRootCmd.Execute())
}

func init() {
	todolistRootCmd.AddCommand(todolistAddCmd, todolistCleanupCmd, todolistDoneCmd, todolistListCmd, todolistUndoneCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// todolistCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// todolistCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	//todolistRootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.toDoList.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//todolistRootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
