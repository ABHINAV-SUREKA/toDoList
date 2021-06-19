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

var filename = "_todolist"

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
		fmt.Println(taskName + " added to the list")
	}
}

func Done(taskId int) {
	var toDoList toDoList
	if fileExists(filename) {
		toDoList = toDoListFromFile(filename)
		toDoList.markDone(taskId)
	} else {
		fmt.Println("No To Do List found\nAdd a task")
	}
}

func UnDone(taskId int) {
	var toDoList toDoList
	if fileExists(filename) {
		toDoList = toDoListFromFile(filename)
		toDoList.markUnDone(taskId)
	} else {
		fmt.Println("No To Do List found\nAdd a task")
	}
}

func List() {
	var toDoList toDoList
	if fileExists(filename) {
		toDoList = toDoListFromFile(filename)
		fmt.Println(toDoList.unDoneTasksToString())
	} else {
		fmt.Println("No To Do List found\nAdd a task")
	}
}

func (tdl toDoList) cleanup() {
	//
}

func Test() {
	fmt.Println("test")
}
