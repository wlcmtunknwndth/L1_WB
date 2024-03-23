package main

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/wlcmtunknwndth/L1_WB/task5/Receiver"
	"github.com/wlcmtunknwndth/L1_WB/task5/Writer"
	"time"
)

// Разработать программу, которая будет последовательно отправлять значения в канал, а
//с другой стороны канала — читать. По истечению N секунд программа должна завершаться.

func main() {
	// Writer
	ch := make(chan *any)
	quit := make(chan bool)

	go func(quit chan bool) {
		time.Sleep(time.Second)
		quit <- true
	}(quit)

	receiver := Receiver.NewReceiver(ch)
	receiver.Process()

	writer := Writer.NewWriter(ch)
	sourceCh := make(chan *any)
	writer.Process(sourceCh)

	for {
		data := any(gofakeit.Emoji())
		select {
		case ch <- &data:
		case <-quit:
			receiver.Stop()
			writer.Stop()

			close(sourceCh)
			close(quit)
			return
		}
	}
}
