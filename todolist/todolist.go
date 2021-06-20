package todolist

import (
	"fmt"
	"os"
)

type task struct {
	taskName string
	done     string
}
type toDoList []task

// flat file to save the task(s)
var filename = "_todolist"

// Add : add a task
func Add(taskName string) {
	var toDoList toDoList
	if fileExists(filename) {
		toDoList = toDoListFromFile(filename)
	} else {
		toDoList = newToDoList()
	}
	toDoList.addTask(taskName)
	err := toDoListToFile(filename, toDoList)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	} else {
		fmt.Println("'" + taskName + "' added to the list")
	}
}

// Done : mark a task as done
func Done(taskId int) {
	var toDoList toDoList
	if fileExists(filename) {
		toDoList = toDoListFromFile(filename)
		toDoList.markDone(filename, taskId)
	} else {
		fmt.Println("No To Do List found\nAdd new task")
	}
}

// UnDone : mark a task as undone
func UnDone(taskId int) {
	var toDoList toDoList
	if fileExists(filename) {
		toDoList = toDoListFromFile(filename)
		toDoList.markUnDone(filename, taskId)
	} else {
		fmt.Println("No To Do List found\nAdd new task")
	}
}

// List : list all undone task(s)
func List() {
	var toDoList toDoList
	if fileExists(filename) {
		toDoList = toDoListFromFile(filename)
		fmt.Println(toDoList.unDoneTasksToString())
	} else {
		fmt.Println("No To Do List found\nAdd new task")
	}
}

// Cleanup : clear all done task(s)
func Cleanup() {
	var toDoList toDoList
	if fileExists(filename) {
		toDoList = toDoListFromFile(filename)
		toDoList.cleanupDoneTasks(filename)
	} else {
		fmt.Println("No To Do List found\nAdd new task")
	}
}
