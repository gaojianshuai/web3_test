package main

import (
	"fmt"
	"math"
)

// 面向对象编程
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle 结构体
type Rectangle struct {
	Width, Height float64
}

// Circle 结构体
type Circle struct {
	Radius float64
}

// Area 矩形面积
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter 矩形周长
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Area 圆面积
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter 圆周长
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func main() {
	rect := Rectangle{Width: 10, Height: 5}
	circ := Circle{Radius: 7}

	fmt.Printf("Rectangle Area: %.2f\n", rect.Area())
	fmt.Printf("Rectangle Perimeter: %.2f\n", rect.Perimeter())

	fmt.Printf("Circle Area: %.2f\n", circ.Area())
	fmt.Printf("Circle Perimeter: %.2f\n", circ.Perimeter())

}
