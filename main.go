package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

// task struct
type Task struct {
	id int
	title string
	completed bool
}


// function DisplayMenu
func DisplayMenu(){
	fmt.Println("Task Manager")

	fmt.Println("1. Add Task")
	fmt.Println("2. View Tasks")
	fmt.Println("3. Complete Task")
	fmt.Println("4. Delete Task")
	fmt.Println("5. Exit")
}


// function AddTask(tasks []Task)
func AddTask(tasks []Task) []Task{
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Your Task Title: ")

	input, err := reader.ReadString('\n')
	if err != nil{
		fmt.Println("An error occurred:", err)
		return tasks
	}

	input = strings.TrimSpace(input)

	nextId := len(tasks) + 1

	newTask := Task{
		id:        nextId,
		title:     input,
		completed: false,
	}

	tasks = append(tasks, newTask)

	return tasks

}


// function ViewTasks(tasks []Task)
func ViewTasks(tasks []Task){
	for _, task := range tasks {
		if task.completed {
			fmt.Printf("[X] %s\n", task.title)
		} else {
			fmt.Printf("[ ] %s\n", task.title)
		}
	}
}


// function CompleteTask(tasks []Task)


// function DeleteTask(task []Task)


func main() {

	tasks := []Task{}
	menuChoice := 0
	reader := bufio.NewReader(os.Stdin)

	for menuChoice != 5 {

		DisplayMenu()


		fmt.Println("Enter your choice.")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occurred:", err)
			continue
		}

		input = strings.TrimSpace(input)

		menuChoice, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("Please enter a number from 1 to 5.")
			continue
		}

		switch menuChoice {
		case 1:
			tasks = AddTask(tasks)
		case 2:
			ViewTasks(tasks)
		case 3:
			CompleteTask(tasks)
		case 4:
			tasks = DeleteTask(tasks)
		case 5:
			fmt.Println("Goodbye!")
		default:
			fmt.Println("Invalid menu option. Pick a number from 1 to 5")
		}

	}
	
}





