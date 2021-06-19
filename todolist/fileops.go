package todolist

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func toDoListFromFile(filename string) toDoList {
	byteSlice, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	toDoList := newToDoList()
	taskInfoList := strings.Split(string(byteSlice), "\n")
	count := 0
	for _, taskInfo := range taskInfoList {
		if len(taskInfo) > 0 {
			count = count + 1
			taskInfoSlice := strings.SplitN(taskInfo, "|", 2)
			task := newTask()
			task.taskName = taskInfoSlice[0]
			task.done = taskInfoSlice[1]
			toDoList = append(toDoList, task)
		}
	}
	return toDoList
}

func toDoListToFile(filename string, tdl toDoList) error {
	return ioutil.WriteFile(filename, []byte(tdl.toString()), 0666)
}
