package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Square struct {
	Length float64
}

type Circle struct {
	Radius float64
}

func getArea(s Shape) float64 {
	return s.Area()
}

func getPerimeter(s Shape) float64 {
	return s.Perimeter()
}

func (s Square) Area() float64 {
	return s.Length * s.Length
}

func (c Circle) Area() float64 {
	return (c.Radius * c.Radius) * math.Pi
}

func (s Square) Perimeter() float64 {
	return s.Length * 4.0
}

func (c Circle) Perimeter() float64 {
	return (c.Radius * 2) * math.Pi
}

func main() {
	var square, circle Shape
	square = Square{Length: 5}
	circle = Circle{Radius: 5}

	fmt.Println("Area of the square -", getArea(square))
	fmt.Println("Area of the circle -", getArea(circle))
	fmt.Println("Perimeter of the square -", getPerimeter(square))
	fmt.Println("Perimeter of the circle -", getPerimeter(circle))
}
