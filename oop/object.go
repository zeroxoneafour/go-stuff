package main

import "fmt"

type Hello struct {
	msg string
}

func (h Hello) sayHello() {
	fmt.Println(h.msg)
}

func main() {
	a_hello := Hello{msg: "Hello World!"}
	a_hello.sayHello()
}
