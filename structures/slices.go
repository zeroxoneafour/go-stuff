package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Numbers 0-4 -", slice[0:5])
	fmt.Println("Numbers 5-9 -", slice[5:10])
}
