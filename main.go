package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
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

	maxID := 0

	for _, task := range tasks{
		if task.id > maxID{
			maxID = task.id
		}
	}

	nextId := maxID + 1

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
func CompleteTask(tasks []Task){

	ViewTasks(tasks)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter a task ID")

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error occurred.")
		return
	}
	
	input = strings.TrimSpace(input)
	
	taskChoice, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("No task ID found")
		return
	}

	found := false

	for i, task := range tasks{
		if task.id == taskChoice{
			tasks[i].completed = true
			found = true
			break

		}
	}
	if !found {
		fmt.Println("Task not found")
}
}


// function DeleteTask(task []Task)
func DeleteTask(tasks []Task) []Task{

	ViewTasks(tasks)
	
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Choose a task ID to delete.")

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An Error occurred.")
		return tasks
	}

	input = strings.TrimSpace(input)

	taskChoice, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("No task ID found")
		return tasks
	}

	found := false

	for i, task := range tasks {
		if task.id == taskChoice {
			tasks = slices.Delete(tasks, i, i+1)
			found = true
			break

		}
		
	}
	if !found {
		fmt.Println("Task not found")
	}

	return  tasks 

}


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





