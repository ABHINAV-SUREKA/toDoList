package todolist

import (
	"fmt"
	"os"
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
	if len(tdl) == 0 {
		return "To Do List empty\nAdd a task"
	}
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

func (tdl *toDoList) markDone(taskId int) {
	if taskId > len(*tdl)-1 || taskId < 0 {
		fmt.Println("Error: invalid task id. Try again")
		os.Exit(1)
	}
	if len(*tdl) == 0 {
		fmt.Println("To Do List empty\nAdd a task")
		return
	}
	for i, task := range *tdl {
		if i == taskId {
			if task.done == "undone" {
				task.done = "done"
				(*tdl)[i] = task
				err := toDoListToFile(filename, *tdl)
				if err != nil {
					fmt.Println("Error:", err)
					os.Exit(1)
				} else {
					fmt.Println(task.taskName + " marked as completed")
					return
				}
			} else {
				fmt.Println(task.taskName + " already done")
				return
			}
		}
	}
}

func (tdl *toDoList) markUnDone(taskId int) {
	if taskId > len(*tdl)-1 || taskId < 0 {
		fmt.Println("Error: invalid task id. Try again")
		os.Exit(1)
	}
	if len(*tdl) == 0 {
		fmt.Println("To Do List empty\nAdd a task")
		return
	}
	for i, task := range *tdl {
		if i == taskId {
			if task.done == "done" {
				task.done = "undone"
				(*tdl)[i] = task
				err := toDoListToFile(filename, *tdl)
				if err != nil {
					fmt.Println("Error:", err)
					os.Exit(1)
				} else {
					fmt.Println(task.taskName + " marked as incomplete")
					return
				}
			} else {
				fmt.Println(task.taskName + " is already incomplete")
				return
			}
		}
	}
}
