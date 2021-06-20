package todolist

import (
	"fmt"
	"os"
	"strings"
)

// newTask : initialise new task
func newTask() task {
	return task{}
}

// newToDoList : initialise new toDoList
func newToDoList() toDoList {
	return toDoList{}
}

// toString : convert all tasks in toDoList to string | excludes a task while conversion if found in delete index
func (tdl toDoList) toString(deleteIndexes ...int) string {
	toDoListStr := ""
	if len(deleteIndexes) != 0 {
		j := 0
		for i, task := range tdl {
			if j < len(deleteIndexes) && i == deleteIndexes[j] {
				j++
				continue
			} else {
				toDoListStr = toDoListStr + task.taskName + "|" + task.done + "\n"
			}
		}
	} else {
		for _, task := range tdl {
			toDoListStr = toDoListStr + task.taskName + "|" + task.done + "\n"
		}
	}
	return toDoListStr
}

// unDoneTasksToString : convert undone task(s) from toDoList to string
func (tdl toDoList) unDoneTasksToString() string {
	toDoListStr := ""
	if len(tdl) == 0 {
		return "To Do List empty\nAdd a task"
	}
	for i, task := range tdl {
		if task.done == "undone" {
			toDoListStr = toDoListStr + fmt.Sprintf("%v %v\n", i+1, task.taskName)
		}
	}
	unDoneTasksStr := strings.TrimSuffix(toDoListStr, "\n")
	if unDoneTasksStr == "" {
		return "Bravo! All tasks completed!\nAdd new task"
	}
	return "Tasks to be done:\n" + unDoneTasksStr
}

// addTask : add a task to toDoList
func (tdl *toDoList) addTask(taskName string) {
	task := task{
		taskName: taskName,
		done:     "undone",
	}
	*tdl = append(*tdl, task)
}

// markDone : mark a task as done
func (tdl *toDoList) markDone(fileName string, taskId int) {
	if taskId > len(*tdl)-1 || taskId < 0 {
		fmt.Println("Error: Invalid task id. Try again")
		os.Exit(1)
	}
	if len(*tdl) == 0 {
		fmt.Println("To Do List empty\nAdd new task")
		return
	}
	for i, task := range *tdl {
		if i == taskId {
			if task.done == "undone" {
				task.done = "done"
				(*tdl)[i] = task
				err := toDoListToFile(fileName, *tdl)
				if err != nil {
					fmt.Println("Error:", err)
					os.Exit(1)
				} else {
					fmt.Println("'" + task.taskName + "' marked as completed")
					return
				}
			} else {
				fmt.Println("'" + task.taskName + "' is already done. Nothing to do")
				return
			}
		}
	}
}

// markUnDone : mark a task as undone
func (tdl *toDoList) markUnDone(fileName string, taskId int) {
	if taskId > len(*tdl)-1 || taskId < 0 {
		fmt.Println("Error: Invalid task id. Try again")
		os.Exit(1)
	}
	if len(*tdl) == 0 {
		fmt.Println("To Do List empty\nAdd new task")
		return
	}
	for i, task := range *tdl {
		if i == taskId {
			if task.done == "done" {
				task.done = "undone"
				(*tdl)[i] = task
				err := toDoListToFile(fileName, *tdl)
				if err != nil {
					fmt.Println("Error:", err)
					os.Exit(1)
				} else {
					fmt.Println("'" + task.taskName + "' marked as incomplete")
					return
				}
			} else {
				fmt.Println("'" + task.taskName + "' has not been done yet. Nothing to do")
				return
			}
		}
	}
}

// cleanupDoneTasks : cleanup all done task(s) from toDoList
func (tdl *toDoList) cleanupDoneTasks(fileName string) {
	if len(*tdl) == 0 {
		fmt.Println("To Do List empty\nAdd new task")
		return
	}
	var deleteIndexes []int
	for i, task := range *tdl {
		if task.done == "done" {
			deleteIndexes = append(deleteIndexes, i)
		}
	}
	if len(deleteIndexes) == 0 {
		fmt.Println("All tasks are incomplete. Nothing to do")
		return
	}
	err := toDoListToFile(fileName, *tdl, deleteIndexes...)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	} else {
		fmt.Println("All the completed tasks cleared from To Do List")
		return
	}
}
