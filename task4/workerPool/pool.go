package workerPool

import (
	"fmt"
	"sync"
	"time"
)

type Pool struct {
	Tasks   []*Task
	Workers []*Worker

	threads       int
	collector     chan *Task
	runBackground chan bool
	wg            sync.WaitGroup
}

func (p *Pool) AddTask(task *Task) {
	p.collector <- task
}

func (p *Pool) Run() {
	go func() {
		for {
			fmt.Printf("No new task arrived...\n")
			time.Sleep(30 * time.Second)
			p.Stop()
		}
	}()

	for i := 1; i <= p.threads; i++ {
		worker := NewWorker(p.collector, i)
		p.Workers = append(p.Workers, worker)
		go worker.Start()
	}

	for i := range p.Tasks {
		p.collector <- p.Tasks[i]
	}

	p.runBackground = make(chan bool)
	<-p.runBackground
}

func (p *Pool) Stop() {
	for i := range p.Workers {
		p.Workers[i].Stop()
	}
	p.runBackground <- true
}

func NewPool(tasks []*Task, workers, maxTasks int) *Pool {
	return &Pool{
		Tasks:     tasks,
		threads:   workers,
		collector: make(chan *Task, maxTasks),
	}
}
