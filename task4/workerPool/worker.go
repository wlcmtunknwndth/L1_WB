package workerPool

import (
	"fmt"
)

type Worker struct {
	ID       int
	taskChan chan *Task
	quit     chan bool
}

func NewWorker(ch chan *Task, ID int) *Worker {
	return &Worker{
		ID:       ID,
		taskChan: ch,
		quit:     make(chan bool),
	}
}

func (wr *Worker) Start() {
	fmt.Printf("Starting worker %d\n", wr.ID)

	for {
		select {
		case task := <-wr.taskChan:
			process(wr.ID, task)

		case <-wr.quit:
			return
		}
	}
}

func (wr *Worker) Stop() {
	fmt.Printf("closing worker %d\n", wr.ID)
	go func() {
		wr.quit <- true
	}()
}
