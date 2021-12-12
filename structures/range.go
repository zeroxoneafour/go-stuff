package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Numbers 0-9")
	for _, j := range slice {
		fmt.Println(j)
	}
}
