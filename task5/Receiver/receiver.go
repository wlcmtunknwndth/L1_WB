package Receiver

import (
	"fmt"
)

type Receiver struct {
	ch   <-chan *any
	quit chan bool
}

func NewReceiver(ch <-chan *any) *Receiver {
	return &Receiver{
		ch:   ch,
		quit: make(chan bool),
	}
}

func (r *Receiver) Process() {
	fmt.Println("Starting receiver")

	go func() {
		for {
			select {
			case data := <-r.ch:
				fmt.Printf("got data: %+v\n", (*data).(string))
			case <-r.quit:
				return
			}
		}
	}()
}

func (r *Receiver) Stop() {
	fmt.Println("closing receiver")
	go func() {
		r.quit <- true
	}()
}
