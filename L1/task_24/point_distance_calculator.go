package main

import (
	"fmt"
	"math"
)

/*
	Разработать программу нахождения расстояния между двумя точками,
	которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
*/

type Point struct {
	x float64
	y float64
}

func (p *Point) distToPoint(point *Point) float64 {
	return math.Sqrt((p.x-point.x)*(p.x-point.x) + (p.y-point.y)*(p.y-point.y))
}

func (p *Point) String() string {
	return fmt.Sprintf("%.1f, %.1f", p.x, p.y)
}

func NewPoint(x float64, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func main() {
	p1 := NewPoint(0, 0)
	p2 := NewPoint(0, 5)

	fmt.Printf("Расстояние между точками (%s) и (%s) = %.1f", p1.String(), p2.String(), p1.distToPoint(p2))

}
