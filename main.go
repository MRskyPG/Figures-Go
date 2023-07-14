package main

import "fmt"

type Figure interface {
	Square() float64
	Length() float64
}

func main() {
	figures := []Figure{
		NewRectangle(5, 15),
		NewCircle(3, 4, 5),
		NewTriangle(4, 5, 6),
	}
	for _, figure := range figures {
		fmt.Println(figure)
	}
	tr := NewTriangle(7, 8, 9)
	fmt.Println(tr.Square())
}
