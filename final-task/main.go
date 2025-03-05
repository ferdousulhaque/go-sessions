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

	fmt.Println("To-Do List:")
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
	} else {
		fmt.Println("Task added!")
	}
	defer file.Close()
}

func deleteTask() {
	fmt.Print("Enter the task number to delete: ")
	var taskNum int
	fmt.Scanln(&taskNum)
	if taskNum > 0 {
		err := DeleteLineInFile(os.Args[1]+".todo", taskNum)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Line deleted successfully.")
		}
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
	fmt.Print("Enter the task number to mark completed: ")
	var taskNum int
	fmt.Scanln(&taskNum)
	if taskNum > 0 {
		err := ModifyLineInFile(os.Args[1]+".todo", taskNum, "[x] ")
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Task marked as completed!")
		}
	}
}

// ModifyLineInFile modifies a specific line in a file by adding a prefix to it.
func ModifyLineInFile(filename string, lineNumber int, prefix string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	currentLine := 1

	for scanner.Scan() {
		line := scanner.Text()
		if currentLine == lineNumber {
			line = prefix + line // Add prefix to the specified line
		}
		lines = append(lines, line)
		currentLine++
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	// Write back to file
	file, err = os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}

// DeleteLineInFile deletes a specific line in a file.
func DeleteLineInFile(filename string, lineNumber int) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	currentLine := 1

	for scanner.Scan() {
		if currentLine != lineNumber {
			lines = append(lines, scanner.Text())
		}
		currentLine++
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	// Write back to file
	file, err = os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}
