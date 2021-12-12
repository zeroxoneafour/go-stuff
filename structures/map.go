package main

import (
	"fmt"
	"sort"
)

func main() {
	numbers := map[int]string{0: "zero", 1: "one", 2: "two", 3: "three", 4: "four", 5: "five", 6: "six", 7: "seven", 8: "eight", 9: "nine"}
	fmt.Println("Printing numbers")
	for i, j := range numbers {
		fmt.Println(i, "-", j)
	}
	fmt.Println("Printing numbers in order")
	keys := make([]int, 0, len(numbers))
	for i := range numbers {
		keys = append(keys, i)
	}
	sort.Ints(keys)
	for i, j := range keys {
		fmt.Println(keys[i], "-", numbers[j])
	}
}
