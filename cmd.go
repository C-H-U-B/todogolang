package main

import (
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

const listPath = "./todo.json"

var rootCmd = &cobra.Command{
	Use: "todo",
	Run: func(cmd *cobra.Command, args []string) {
		_ = cmd.Help()
	},
}

var addCmd = &cobra.Command{
	Use: "add",
	Run: func(cmd *cobra.Command, args []string) {
		todo, err := LoadTodoList(listPath)
		if err != nil {
			todo = TodoList{}
		}
		item := TodoItem{
			Description: args[0],
		}
		todo.Add(item)
		todo.Save(listPath)
	},
}


var displayCmd = &cobra.Command{
	Use: "display",
	Run: func(cmd *cobra.Command, args []string) {
		todo, err := LoadTodoList(listPath)
		if err != nil {
			todo = TodoList{}
		}
		for i, item := range todo {
			println("la tache ", i, " ", item.Description, " est à faire avant le ", item.Date.String())
		}
	},
}

var deleteCmd = &cobra.Command{
	Use: "delete",
	Run: func(cmd *cobra.Command, args []string) {
		todo, err := LoadTodoList(listPath)
		if err != nil {
			todo = TodoList{}
		}
		index, err := strconv.Atoi(args[0])
		if err != nil {
			panic("invalid task id")
		}
		todo.Delete(index)
		todo.Save(listPath)
	},
}


func Run() {
	rootCmd.AddCommand(addCmd)
	rootCmd.AddCommand(displayCmd)
	rootCmd.AddCommand(deleteCmd)
	if err := rootCmd.Execute(); err != nil {
		_ = rootCmd.Help()
		os.Exit(1)
	}
}
