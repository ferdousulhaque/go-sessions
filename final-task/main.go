package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type option struct {
	Text  string
	Value int
}

func main() {
	if os.Args[1] != "" {
		if _, err := os.Stat(os.Args[1]); os.IsNotExist(err) {
			_, e := os.Create(os.Args[1])
			if e != nil {
				log.Fatal(e)
			}
		}
	} else {
		fmt.Println("Please provide a file path")
		os.Exit(1)
	}

	fmt.Println("Welcome to your daily To-Do list!")
	fmt.Println("Please select an option:")
	options := []option{
		{Text: "Add a task", Value: 1},
		{Text: "Delete a task", Value: 2},
		{Text: "View tasks", Value: 3},
		{Text: "Mark task as completed", Value: 4},
		{Text: "Exit", Value: 5},
	}
	// Print options
	for _, opt := range options {
		fmt.Printf("%d. %s\n", opt.Value, opt.Text)
	}

	// Get user input
	var choice int
	fmt.Print("Enter your choice: ")
	fmt.Scanln(&choice)

	// Process user input
	switch choice {
	case 1:
		addTask()
	case 2:
		deleteTask()
	case 3:
		viewTasks()
	case 4:
		markTaskAsCompleted()
	case 5:
		os.Exit(0)
	default:
		fmt.Println("Invalid choice. Please try again.")
	}
}

func addTask() {
	fmt.Println("Adding a task...")
}

func deleteTask() {
	fmt.Println("Deleting a task...")
}

func viewTasks() {
	//read from a file
	filePath := os.Args[1]
	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	for _, line := range fileLines {
		fmt.Println(line)
	}

	fmt.Println(fileLines)
}

func markTaskAsCompleted() {
	fmt.Println("Marking a task as completed...")
}
