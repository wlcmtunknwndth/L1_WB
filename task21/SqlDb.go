package main

import "fmt"

type SqlDb struct {
}

type Data struct {
}

func (s *SqlDb) SaveData(d *Data) {
	fmt.Println("Saved your data to SQL DB")
}

func (s *SqlDb) GetData(name string) {
	fmt.Println("Got your order from SQL DB")
}
