package todolist

import (
	"fmt"
	"io/ioutil"
	"os"
)

func newToDoList() toDoList {
	return toDoList{}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func toDoListFromFile(filename string) {
	_, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}

func (tdl toDoList) toString() string {
	toDoListStr := ""
	for _, task := range tdl {
		toDoListStr = toDoListStr + task.taskName + "\t" + task.done + "\n"
	}
	return toDoListStr
}
