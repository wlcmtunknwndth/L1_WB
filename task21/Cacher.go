package main

import "fmt"

type Cacher struct {
}

func (s *Cacher) Cache(d *Data) {
	fmt.Println("Saved your data to NoSQL")
}

func (s *Cacher) GetCache(name string) {
	fmt.Println("Got your data from NoSQL")
}
