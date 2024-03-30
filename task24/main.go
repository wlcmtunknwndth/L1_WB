package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func (p *Point) Dist() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

func main() {
	p1 := Point{
		x: 12314.124,
		y: 241.4121,
	}

	fmt.Println(p1.Dist())
}
