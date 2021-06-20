package todolist

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// fileExists : check if the file exist
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

// toDoListFromFile : read data from the file and convert it to toDoList
func toDoListFromFile(filename string) toDoList {
	byteSlice, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	toDoList := newToDoList()
	taskInfoList := strings.Split(string(byteSlice), "\n")
	for _, taskInfo := range taskInfoList {
		if len(taskInfo) > 0 {
			taskInfoSlice := strings.SplitN(taskInfo, "|", 2)
			task := newTask()
			task.taskName = taskInfoSlice[0]
			task.done = taskInfoSlice[1]
			toDoList = append(toDoList, task)
		}
	}
	return toDoList
}

// toDoListToFile : convert toDoList to string and write to the file | ignore done task(s) to be cleared, if any
func toDoListToFile(filename string, tdl toDoList, deleteIndexes ...int) error {
	return ioutil.WriteFile(filename, []byte(tdl.toString(deleteIndexes...)), 0666)
}
