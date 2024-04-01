// In Go (Golang), a function variable is a variable that can hold a reference to a function. This allows you to treat functions as first-class citizens, meaning you can pass them as arguments to other functions, return them from other functions, and assign them to variables.

// Here's an example to illustrate function variables in Go:

package main

import "fmt"

// Function type signature
type MathFunc func(int, int) int

// Add function
func add(a, b int) int {
	return a + b
}

// Subtract function
func subtract(a, b int) int {
	return a - b
}

func main() {
	var operation MathFunc // Declare a function variable of type MathFunc

	operation = add // Assign the add function to the operation variable
	result := operation(10, 5)
	fmt.Println("Addition result:", result) // Output: Addition result: 15

	operation = subtract // Assign the subtract function to the operation variable
	result = operation(10, 5)
	fmt.Println("Subtraction result:", result) // Output: Subtraction result: 5

	fmt.Println("======================================")






    

	// Immediately invoked anonymous function
	func() {
		fmt.Println("Inside IIFE-like function")
	}() // Immediately invoke the anonymous function

	// You can also use variables inside the IIFE-like function
	message := "Hello from IIFE-like function"
	func(msg string) {
		fmt.Println(msg)
	}(message) // Pass and immediately invoke the anonymous function with the message
}

// In this example:

// MathFunc is a type that represents a function that takes two int parameters and returns an int.
// add and subtract are two functions that match the MathFunc type signature.
// We declare a variable operation of type MathFunc. This variable can hold references to functions that match the MathFunc signature.
// We assign the add function to operation, and then we call operation with arguments 10 and 5, which executes the add function.
// Next, we assign the subtract function to operation, and then we call operation again with arguments 10 and 5, which executes the subtract function.
// Function variables are useful when you want to change the behavior of a function dynamically or when you want to create higher-order functions that can operate on different functions based on some conditions or inputs.
