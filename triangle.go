package main

import (
	"errors"
	"fmt"
	"math"
)

type Triangle struct {
	A, B, C int
}

func NewTriangle(A, B, C int) Triangle {
	t, err := isTriangle(Triangle{A, B, C})
	if err != nil {
		fmt.Println("An error occurred: ", err)
		return Triangle{}
	}
	return t
}

func isTriangle(t Triangle) (Triangle, error) {
	if t.A < t.B+t.C && t.B < t.A+t.C && t.C < t.A+t.B {
		return t, nil
	}
	return Triangle{}, errors.New("Triangle with these sides doesn't exist")
}

func (t Triangle) String() string {
	return fmt.Sprintf("Triangle. Sides: %d, %d and %d", t.A, t.B, t.C)
}

func (t Triangle) Square() float64 {
	p := 0.5 * t.Length()
	height := float64(2.0) / float64(t.A) * math.Sqrt(p*(p-float64(t.A))*(p-float64(t.B))*(p-float64(t.C)))
	return 0.5 * height * float64(t.A)
}

func (t Triangle) Length() float64 {
	return float64(t.A + t.B + t.C)
}
