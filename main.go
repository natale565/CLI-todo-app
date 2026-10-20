package main

import (
	"bufio"
	"fmt"
	"os"
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

}





