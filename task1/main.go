package main

import (
	"fmt"
	"strings"
)

type Human struct {
	fio string
	age uint8
}

func (h *Human) getFIO() string {
	return h.fio
}

func (h *Human) getAge() uint8 {
	return h.age
}

func (h *Human) setFIO(name, surname string) bool {
	if name != "" && surname != "" {
		h.fio = strings.Join([]string{name, surname}, " ")
		return true
	}
	return false
}

func (h *Human) setAge(age uint8) bool {
	if age > 150 {
		return false
	}
	h.age = age
	return true
}

type Action struct {
	Human
}

func main() {
	var act Action
	act.setAge(22)
	act.setFIO("Artyem", "Petrov")

	fmt.Printf("%+v", act)
}
