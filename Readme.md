# What is this?
Basic task manager command-line application using Cobra to learn more about CLI development in Go lang

# Requirements
## 1. commands

- add: Add a new task.
- complete: Mark a task as completed.
- list: List all tasks.
- remove: Remove a task.

## 2. flags and argument
- add command should take a task description as an argument.
- complete and remove commands should take the task index as an argument.
- You can use flags to provide additional options, such as a due date or priority level.

## 3. Task storage

- Store tasks in-memory or use a simple JSON file
- possibility to handle configuration and set the storage location.

## 4. User friendly?

- Use different colors for different types of output (success, error, etc.).
- Provide clear and user-friendly output for each command.

## 5. Error handling 
- Implement proper error handling for cases like invalid input, out-of-bounds task index

## 6. Testing
- Write unit tests for the core functionality of your task manager
