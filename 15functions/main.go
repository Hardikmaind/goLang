package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()

	result := adder(3, 5)
	fmt.Println("Result is: ", result)

	proRes, myMessage := proAdder(2, 5, 8, 7, 3)
	fmt.Println("Pro result is: ", proRes)
	fmt.Println("Pro Message is: ", myMessage)


	
	fmt.Println("=====================================")
	message := "Hello from main"
	
	// Defining a function variable that accesses a variable from the parent function
	nestedFunc := func() {
		fmt.Println(message)
	}
	
	// Calling the function variable
	nestedFunc() // Output: Hello from main

}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

// in go we can return multiple values from a function
//in below function we are returning two values one is the total and other is the message
func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}

	return total, "Hi Pro result function"

}

func greeter() {
	fmt.Println("Namstey from golang")
}
