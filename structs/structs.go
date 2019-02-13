package structs

import "math"

// Perimeter return perimeter of rectangle
// Requirement:
// Calculate permimeter of a rectangle
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Height + rectangle.Width)
}

// Area returns area of rectangle
// Requirement:
// Calculate area of a rectangle

// Rectangle has width and height
type Rectangle struct {
	Width  float64
	Height float64
}

// Area if rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle has radius
type Circle struct {
	Radius float64
}

// Area of a circle
func (c Circle) Area() float64 {
	return (c.Radius * c.Radius * math.Pi)
}

// Triangle has different area
type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return t.Base * t.Height * 0.5
}

// Shape is an interface to implements Area()
type Shape interface {
	Area() float64
}
