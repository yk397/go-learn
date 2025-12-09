package oop

import (
	"fmt"
	"math"
)

// 题目一
type Shape interface {
	Area()
	Perimeter()
}

type Rectangle struct {
	Length float64
	Width  float64
}

type Circle struct {
	Radius float64
}

func (rec *Rectangle) Area() float64 {
	return rec.Length * rec.Width
}

func (rec *Rectangle) Perimeter() float64 {
	return (rec.Length + rec.Width) * 2
}

func (c *Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// 题目二
type Person struct {
	Name string
	Age  int8
}

type Employee struct {
	Person
	EmployeeId int32
}

func (e Employee) PrintInfo() {
	fmt.Println(e)
}
