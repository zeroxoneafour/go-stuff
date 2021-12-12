package main

import "fmt"

func fibonacci() func() int {
	num1 := 0
	num2 := 1
	return func() int {
		ret := num1
		num1 += num2
		num2 = ret
		return ret
	}
}

func main() {
	f := fibonacci()
	var nums [20]int
	for i := 0; i < len(nums); i++ {
		nums[i] = f()
	}
	fmt.Println("fibonacci", nums)
}
