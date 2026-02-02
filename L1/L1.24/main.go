package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

func (p *Point) Distance(q *Point) float64 {
	return math.Hypot(p.x-q.x, p.y-q.y)
}

func main() {
	a := NewPoint(0, 0)
	b := NewPoint(1, 0)
	fmt.Println(a.Distance(b))

	a = NewPoint(1, 2)
	b = NewPoint(4, 6)
	fmt.Println(a.Distance(b))
}