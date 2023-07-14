package main

import (
	"fmt"
	"math"
)

// Circle with central point (X,Y) and Radius
type Circle struct {
	X      int
	Y      int
	Radius int
}

func NewCircle(X, Y, Radius int) Circle {
	return Circle{X, Y, Radius}
}

func (c Circle) String() string {
	return fmt.Sprintf("Circle. Central point = (%d, %d), Radius = %d", c.X, c.Y, c.Radius)
}

func (c Circle) Square() float64 {
	return math.Pi * float64(c.Radius*c.Radius)
}

func (c Circle) Length() float64 {
	return 2 * math.Pi * float64(c.Radius)
}
