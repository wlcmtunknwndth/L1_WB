package main

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"sync"
	"time"
)

type ConcurrentMap struct {
	mtx  sync.RWMutex
	data map[string]string
}

func (s *ConcurrentMap) ReadLock() { s.mtx.RLock() }

func (s *ConcurrentMap) ReadUnlock() { s.mtx.RUnlock() }

func (s *ConcurrentMap) WriteLock() { s.mtx.Lock() }

func (s *ConcurrentMap) WriteUnlock() { s.mtx.Unlock() }

func (s *ConcurrentMap) Save(key, value string) {
	s.WriteLock()
	defer s.WriteUnlock()
	s.data[key] = value
}

func (s *ConcurrentMap) Get(key string) (string, error) {
	defer s.ReadUnlock()
	s.ReadLock()
	if data, ok := s.data[key]; ok {
		return data, nil
	}
	return "", fmt.Errorf("couldn't find value by key %s\n", key)
}

func NewMap(len uint8) *ConcurrentMap {
	return &ConcurrentMap{
		data: make(map[string]string, len),
	}
}

func main() {
	start := time.Now()
	var names = []string{"artyem", "liza", "katya", "jenya", "senya", "arina", "vlad", "nikolay petrovich", "alexandra", "murad"}
	var wg sync.WaitGroup
	phoneBook := NewMap(uint8(len(names)))

	// fill with random values
	for i := range names {
		wg.Add(1)
		go func(name string) {
			defer wg.Done()
			rnd := gofakeit.Phone()
			phoneBook.Save(name, rnd)
			fmt.Printf("succesfully saved key %s with value %s\n", name, rnd)
		}(names[i])
	}
	wg.Wait()

	for i := range names {
		wg.Add(1)
		go func(name string) {
			defer wg.Done()
			phone, err := phoneBook.Get(name)
			if err != nil {
				fmt.Printf("couldn't get data: key %s: %v\n", name, err)
				return
			}
			fmt.Printf("received number of %s: %s\n", name, phone)
		}(names[i])
	}

	wg.Wait()
	fmt.Println(time.Since(start))
}
