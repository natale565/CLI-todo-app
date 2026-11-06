# CLI Task Manager

A simple command-line task manager built with Go.

I created this project while learning Go to reinforce the concepts I was studying by applying them in a small, practical project. Rather than continuing to only follow tutorials and lessons, I wanted to build something independently and work through the problems that came up along the way.

## Features

The application currently allows users to:

- Add new tasks
- View all tasks
- Mark tasks as completed
- Delete tasks
- Assign unique IDs to tasks
- Validate user input
- Prevent empty task titles
- Handle invalid task IDs
- Handle an empty task list

Tasks are displayed with their completion status:

```text
1. [ ] Learn Go
2. [X] Build CLI task manager
3. [ ] Practice working with JSON
```

## Technologies

- Go
- Go Standard Library
- `bufio`
- `fmt`
- `os`
- `slices`
- `strconv`
- `strings`

## Concepts Practiced

This project helped reinforce several fundamental Go concepts, including:

- Structs
- Slices
- Functions
- Loops
- Conditional statements
- Switch statements
- User input
- Error handling
- String manipulation
- Type conversion
- Working with Go modules
- Adding and removing values from slices

## Running the Project

Clone the repository and navigate into the project directory.

Run the application directly with:

```bash
go run .
```

You can also compile the project:

```bash
go build -o todo
```

Then run the compiled executable:

```bash
./todo
```

## Menu

When the application starts, the user is presented with the following menu:

```text
Task Manager

1. Add Task
2. View Tasks
3. Complete Task
4. Delete Task
5. Exit
```

The menu continues running until the user selects the exit option.

## Task IDs

Each task is assigned a unique numeric ID.

When creating a new task, the application searches the existing task list for the highest ID and assigns the next available number. This prevents duplicate IDs when tasks have previously been deleted.

## Project Goals

The main goal of this project was to become more comfortable writing Go without relying on step-by-step tutorials.

Building the application required breaking the project into smaller problems, researching unfamiliar concepts, debugging issues, and deciding how the program should be structured.

## Planned Improvements

The next major feature is task persistence using JSON.

Planned improvements include:

- Save tasks to a JSON file
- Load previously saved tasks when the application starts
- Preserve completed tasks between sessions
- Continue improving input validation and error handling
- Refactor repeated logic where appropriate

## What I Learned

One of the biggest takeaways from this project was the value of learning through building.

Tutorials and lessons are useful for introducing new concepts, but working through a project independently helped reinforce those concepts by requiring me to apply them to actual problems.

Even a small project like this provided practice with program structure, debugging, error handling, data management, and thinking through how different parts of an application should work together.

## Author

Christopher Natale