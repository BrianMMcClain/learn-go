package main 

import "fmt"

func swap(x string, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("Hello", "World")
	fmt.Println("Hello World swapped is", a, b)
}