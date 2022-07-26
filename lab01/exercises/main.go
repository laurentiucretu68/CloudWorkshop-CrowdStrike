package main

import (
	"example.com/exercises/rick"
	"fmt"
	"math/rand"
	"strings"
)

/// Ex2
func generateNumbers(size int) []int {
	var x = make([]int, size)
	for i := 0; i < size; i++ {
		x[i] = rand.Intn(100)
	}
	return x
}

func addFunc(X, size int, x []int) []int {
	for i := 0; i < size; i++ {
		x[i] = X ^ 2 + 3*X + 5
	}
	return x
}

/// Ex3
func f(x, y int, s1, s2 string) string {
	arr := strings.Index(s1, s2)
	newArr := s1[arr:]
	return newArr[x:(y + 1)]

}

func main() {
	fmt.Println("Hello from main")
	rick.MainFunc()

	// Ex2
	size := 10
	X := 5
	x := generateNumbers(size)
	fmt.Println(x)
	x = addFunc(X, size, x)
	fmt.Println(x)

	// Ex3
	fmt.Println(f(1, 3, "moduleodle", "od"))

	// Ex 5

}
