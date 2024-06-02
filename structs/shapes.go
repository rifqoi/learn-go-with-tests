package structs

import (
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	width  float64
	height float64
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type Square struct {
	width  float64
	height float64
}

func (s Square) Area() float64 {
	return s.height * s.width
}

func Perimeter(rect Rectangle) float64 {
	return 2 * (rect.width + rect.height)
}
