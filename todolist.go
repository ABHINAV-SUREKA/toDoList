package main

type todolist struct {
	index int
	task string
}

func (tdl todolist) add(taskName string) {
	//
}

func (tdl todolist) done(taskId int) {
	//
}

func (tdl todolist) undone(taskId int) {
	//
}

func (tdl todolist) list() []todolist {
	return nil
}

func (tdl todolist) cleanup() {
	//
}