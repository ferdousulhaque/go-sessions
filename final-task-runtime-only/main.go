package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

type option struct {
	Text  string
	Value int
}

var taskList []string

func main() {
	fmt.Println("To-Do App:")
	options := []option{
		{Text: "Add a Task", Value: 1},
		{Text: "List tasks", Value: 2},
		{Text: "Mark a Task as completed", Value: 3},
		{Text: "Delete a Task", Value: 4},
		{Text: "Exit", Value: 5},
	}
	// Print options
	for _, opt := range options {
		fmt.Printf("%d. %s\n", opt.Value, opt.Text)
	}

	// Get user input
	var choice int
	for {
		fmt.Print("Enter choice: ")
		fmt.Scanln(&choice)

		// Process user input
		switch choice {
		case 1:
			addTask()
		case 2:
			viewTasks()
		case 3:
			markTaskAsCompleted()
		case 4:
			deleteTask()
		case 5:
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please try again.")
		}

		// Print options again
		fmt.Println("\nTo-Do App:")
		for _, opt := range options {
			fmt.Printf("%d. %s\n", opt.Value, opt.Text)
		}
	}
}

func addTask() {
	var task string
	fmt.Print("Enter Task:")
	reader := bufio.NewReader(os.Stdin)
	task, _ = reader.ReadString('\n')
	taskList = append(taskList, task)
	fmt.Println("Task added successfully.")
}

func deleteTask() {
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Invalid task number.")
			}
		}()

		fmt.Print("Enter the task number to delete: ")
		var taskNum int
		fmt.Scanln(&taskNum)

		if taskNum <= 0 || taskNum > len(taskList) {
			fmt.Println("Invalid task number.")
			return
		}

		taskList = slices.Delete(taskList, taskNum-1, taskNum)
		fmt.Println("Task deleted successfully.")
	}()
}

func viewTasks() {
	fmt.Println("\nYour Tasks:")

	fmt.Println("")
	for i, task := range taskList {
		fmt.Printf("%d. %s", i+1, task)
	}
	fmt.Println("")
}

func markTaskAsCompleted() {
	fmt.Print("Enter the task number to mark completed: ")
	var taskNum int
	fmt.Scanln(&taskNum)
	if taskNum <= 0 || taskNum > len(taskList) {
		fmt.Println("Invalid task number.")
		return
	}
	taskList[taskNum-1] = "[X] " + taskList[taskNum-1]
	fmt.Println("Task marked as completed.")
}
