package main

import (
	"example.com/exercises/ex1"
	"example.com/exercises/ex2"
	"example.com/exercises/ex3"
	"fmt"
)

func main() {
	// ex1
	fmt.Println(ex1.WordCount("test are multe are test mere"))

	// ex2
	s := []string{"ana", "are", "peree"}
	fmt.Println(ex2.VowelsCount(s))

	// ex3
	var rectangle = ex3.Rectangle{
		Width:  10,
		Length: 20,
	}

	square := ex3.Square{
		Side: 30,
	}

	circle := ex3.Circle{
		Radius: 10,
	}

	fmt.Println(rectangle.VisitForRectangle())
	fmt.Println(square)
	fmt.Println(circle)
}
