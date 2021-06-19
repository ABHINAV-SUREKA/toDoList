package todolist

import (
	"fmt"
	"strings"
)

func newTask() task {
	return task{}
}

func newToDoList() toDoList {
	return toDoList{}
}

func (tdl toDoList) toString() string {
	toDoListStr := ""
	for _, task := range tdl {
		toDoListStr = toDoListStr + task.taskName + "|" + task.done + "\n"
	}
	return toDoListStr
}

func (tdl toDoList) unDoneTasksToString() string {
	toDoListStr := ""
	fmt.Println("Tasks to be done:")
	for i, task := range tdl {
		if task.done == "undone" {
			toDoListStr = toDoListStr + fmt.Sprintf("%v %v\n", i, task.taskName)
		}
	}
	return strings.TrimSuffix(toDoListStr, "\n")
}

func (tdl *toDoList) addTask(taskName string) {
	task := task{
		taskName: taskName,
		done:     "undone",
	}
	*tdl = append(*tdl, task)
}
