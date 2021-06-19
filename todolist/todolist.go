package todolist

import (
	"fmt"
	"io/ioutil"
	"os"
)

type task struct {
	taskName string
	done     string
}
type toDoList []task

var filename = "_todolist"

func Add(taskName string) {
	if fileExists(filename) {
		toDoListFromFile(filename)
	} else {
		toDoList := newToDoList()
		task := task{
			taskName: taskName,
			done:     "undone",
		}
		toDoList = append(toDoList, task)

		err := ioutil.WriteFile(filename, []byte(toDoList.toString()), 0666)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		} else {
			fmt.Println(taskName + " added to the list")
		}
	}
}

func (tdl toDoList) done(taskId int) {
	//
}

func (tdl toDoList) undone(taskId int) {
	//
}

func (tdl toDoList) list() []toDoList {
	return nil
}

func (tdl toDoList) cleanup() {
	//
}

func Test() {
	fmt.Println("test")
}
