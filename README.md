# toDoList
##### To Do List CLI is a [Golang](https://golang.org/) CLI, created using [Cobra](https://github.com/spf13/cobra), to help you organise your daily tasks with following functionalities:
1. `todolist add [task_name]` - add a task
2. `todolist list` - list all incomplete tasks
3. `todolist cleanup` - clear all completed tasks
4. `todolist done [task_id]` - mark a task as complete
5. `todolist undone [task_id]` - mark a task as in complete

### Installation and use
Run the following command in your terminal: <br>
```
go install github.com/ABHINAV-SUREKA/toDoList
todolist --help
```
Output:
```
To Do List CLI is a tool to organise your daily tasks with add, remove, done and undone functionalities

Usage:
  todolist [command]

Available Commands:
  add         Add task to the list
  cleanup     Cleanup done tasks
  done        Mark task as done
  help        Help about any command
  list        List all tasks still to do
  undone      Mark task as not done

Flags:
  -h, --help   help for todolist

Use "todolist [command] --help" for more information about a command.
```

### Alternate way
In case the above installation doesn't work, run the `./main` executable file as follows:
* `./main --help`
* `./main add [task_name]` - add a task
* `./main list` - list all incomplete tasks
* `./main cleanup` - clear all completed tasks
* `./main done [task_id]` - mark a task as complete
* `./main undone [task_id]` - mark a task as in complete
