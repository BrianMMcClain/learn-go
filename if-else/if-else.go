package main 

import "fmt"

func evenodd(x int) string {
	if x % 2 == 0 {
		return "even"
	} else {
		return "odd"
	}

	return "ERROR"
}

func main() {
	fmt.Println("177 is", evenodd(177))
}