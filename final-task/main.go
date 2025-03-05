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
		filename := os.Args[1] + ".todo"
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			_, e := os.Create(filename)
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
	todo := ""
	fmt.Println("Enter Task:")
	reader := bufio.NewReader(os.Stdin)
	todo, _ = reader.ReadString('\n')
	filePath := os.Args[1] + ".todo"
	file,
		err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
	}
	_, err = file.WriteString("\n" + todo)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
}

func deleteTask() {
	fmt.Print("Enter the task number to delete: ")
	var taskNum int
	fmt.Scanln(&taskNum)
	if taskNum > 0 {
		fmt.Printf("Task ID %d deleted successfully\n", taskNum)
	}

}

func viewTasks() {
	//read from a file
	filePath := os.Args[1] + ".todo"
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

	fmt.Println("\nYour Tasks:")
	for index, line := range fileLines {
		if line == "" {
			continue
		}
		fmt.Printf("%d. %s\n", index+1, line)
	}
	fmt.Println("")
	// fmt.Println(fileLines)
}

func markTaskAsCompleted() {
	fmt.Println("Marking a task as completed...")
}
