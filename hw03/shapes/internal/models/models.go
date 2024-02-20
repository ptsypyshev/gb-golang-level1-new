package models

import "math"

type Shape struct{}

func (s *Shape) Area() float64 {
	return 0
}

type Rectangle struct {
	Shape
	sideA float64
	sideB float64
}

func NewRectangle(a, b float64) *Rectangle {
	return &Rectangle{
		sideA: a,
		sideB: b,
	}
}

func (r *Rectangle) Area() float64 {
	return r.sideA * r.sideB
}

type Circle struct {
	Shape
	radius float64
}

func NewCircle(r float64) *Circle {
	return &Circle{
		radius: r,
	}
}

func (c *Circle) Area() float64 {
	return 2 * math.Pi * c.radius * c.radius
}
