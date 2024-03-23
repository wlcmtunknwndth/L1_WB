package Writer

import (
	"fmt"
)

type Writer struct {
	ch   chan<- *any
	quit chan bool
}

func NewWriter(ch chan<- *any) *Writer {
	return &Writer{
		ch:   ch,
		quit: make(chan bool),
	}
}

func (w *Writer) Process(ch chan *any) {
	fmt.Println("Starting writer")

	go func() {
		for {
			data := <-ch
			select {
			case w.ch <- data:
				fmt.Printf("sent data %v\n", (*data).(string))
			case <-w.quit:
				return
			}
		}
	}()
}

func (w *Writer) Stop() {
	fmt.Println("closing writer")
	go func() {
		w.quit <- true
	}()
}
