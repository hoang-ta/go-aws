package main

import (
	"fmt"
)

const (
	PI = 3.14
)

type Circle struct {
	radius float64
}

func (circle Circle) circumference() float64 {
	return circle.radius * 2 * PI
}

func (circle Circle) area() float64 {
	return circle.radius * circle.radius * PI
}

func NewCircle(radius float64) Circle {
	return Circle{
		radius: radius,
	}
}

func main() {
	fmt.Println("Circle")
	circle := NewCircle(3)
	fmt.Println(circle.circumference())
	fmt.Println(circle.area())
}