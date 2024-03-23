package main

import (
	"fmt"
	"github.com/wlcmtunknwndth/L1_WB/task4/workerPool"
	"math/rand"
	"time"
)

func main() {
	var allTasks []*workerPool.Task
	pool := workerPool.NewPool(allTasks, 5, 100)
	go func() {
		for {
			taskID := rand.Intn(100)
			//var command string
			//if _, err := fmt.Scan(&command); err != nil{
			//	fmt.Println("couldn't scan command")
			//}
			//
			time.Sleep(time.Duration(rand.Intn(5)) * time.Second)

			task := workerPool.NewTask(func(data any) error {
				taskID := data.(int)
				time.Sleep(100 * time.Millisecond)
				fmt.Printf("Task %d processed\n", taskID)
				return nil
			}, taskID)
			pool.AddTask(task)
		}
	}()

	pool.Run()
}
