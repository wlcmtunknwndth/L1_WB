package randomChannel

import "github.com/brianvoe/gofakeit/v6"

func FillCh() chan string {
	ch := make(chan string, 10)
	for i := 0; i < cap(ch); i++ {
		ch <- gofakeit.Emoji()
	}
	return ch
}
