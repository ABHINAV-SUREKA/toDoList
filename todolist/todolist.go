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
	//
}

func (tdl toDoList) undone(taskId int) {
	//
}

func List() {
	var toDoList toDoList
	if fileExists(filename) {
		toDoList = toDoListFromFile(filename)
		if len(toDoList) > 0 {
			fmt.Println(toDoList.unDoneTasksToString())
		} else {
			fmt.Println("You have finished all tasks!")
		}
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
