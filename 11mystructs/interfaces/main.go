package main

import (
	"fmt"
	"math"
)

// shape interface
type shape interface {
	area() float64
	circumf() float64
}

type square struct {
	length float64
}
type circle struct {
	radius float64
}

// square methods
func (s square) area() float64 {
	return s.length * s.length
}
func (s square) circumf() float64 {
	return s.length * 4
}

// circle methods
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) circumf() float64 {
	return 2 * math.Pi * c.radius
}

// if there were no interfaces we will have to create a function for each type...like below we have created for cirecle and the commented part is for square
func printShapeInfo2(s circle) {
	fmt.Printf("area of %T is: %0.2f \n", s, s.area())
	fmt.Printf("circumference of %T is: %0.2f \n", s, s.circumf())
}

// func printShapeInfo2(s square) {
// 	fmt.Printf("area of %T is: %0.2f \n", s, s.area())
// 	fmt.Printf("circumference of %T is: %0.2f \n", s, s.circumf())
// }

//but since we have created an interface we can use the same function for both circle and square...see below main function for example

// so basically interfaces are used to create a common method for different types of structs. and to use the same method for different types of structs
//interfaecs are used to group together different types of structs based on the methods they have in common

func printShapeInfo(s shape) {
	fmt.Printf("area of %T is: %0.2f \n", s, s.area())
	fmt.Printf("circumference of %T is: %0.2f \n", s, s.circumf())
}

func main() {
	shapes := []shape{
		square{length: 15.2},
		circle{radius: 7.5},
		circle{radius: 12.3},
		square{length: 4.9},
	}

	for _, v := range shapes {
		printShapeInfo(v)
		fmt.Println("---")
	}

	fmt.Println("=====================================================================")
	printShapeInfo2(circle{radius: 7.5})
	// printShapeInfo2(square{length: 15.2})
}
