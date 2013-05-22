package main

import "fmt"

func swap(x, y string) (a, b string) {
    a = y
    b = x
    return
}

func main() {
	x, y := swap ("Hello", "World")
    fmt.Println("Hello World swapped is", x, y)
}
