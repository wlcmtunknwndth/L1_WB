package workerPool

import "fmt"

type Task struct {
	Err     error
	Data    any
	handler func(any) error
}

func NewTask(f func(any) error, data any) *Task {
	return &Task{handler: f, Data: data}
}

func process(workerID int, task *Task) {
	fmt.Printf("Worker %d processes task %v\n", workerID, task.Data)
	task.Err = task.handler(task.Data)
}
