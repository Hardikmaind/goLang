package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class on pointers")

	// var ptr *int
	// fmt.Println("Value of pointer is ", ptr)

	myNumber := 23

	var ptr = &myNumber

	fmt.Println("Value of actual pointer is ", ptr)
	fmt.Println("Value of actual pointer is ", *ptr)

	*ptr = *ptr + 2
	fmt.Println("New value is: ", myNumber)

	var ptr2 *int = &myNumber
	fmt.Println("Value of actual pointer is ", ptr2)

	fmt.Print("=====================\n")

	var num int = 42
	fmt.Println("Value of num Before:", num) // Before: 42
	modifyValue(&num)                        // Pass the address of 'num' to the function
	fmt.Println("Value of num  After:", num) // After: 100

	fmt.Print("=====================\n")

	fmt.Println("Below is the function to swap two numbers")
	a := 2
	b := 7
	swap(&a, &b)
	fmt.Println("Value of a after swap is ", a)
	fmt.Println("Value of b after swap is ", b)

	fmt.Println("lets try swapping two numbers without swapping function")
	x:=100
	y:=200
	swapWithoutpointer(x, y)
	fmt.Println("Value of x after swap is ", x)
	fmt.Println("Value of y after swap is ", y)


	fmt.Println("=============================================")

	 // Initialize two numbers
	 num1 := 10
	 num2 := 20
 
	 fmt.Println("Before swapping:")
	 fmt.Println("num1 =", num1)
	 fmt.Println("num2 =", num2)
 
	 // Swap numbers using reference
	 num1, num2 = swapByReference(num1, num2)
 
	 fmt.Println("After swapping:")
	 fmt.Println("num1 =", num1)
	 fmt.Println("num2 =", num2)
}

func modifyValue(ptr *int) {
	*ptr = 100 // Modify the value at the memory address pointed by 'ptr'
}

// this is the function to swap two numbers
func swap(x *int, y *int) {
	temp := *x
	*x = *y
	*y = temp
}
// In this function, you're trying to swap the values of x and y. However, since Go is pass-by-value (meaning it creates a copy of the arguments passed to a function), the changes you make to x and y inside the function don't affect the original variables.
func swapWithoutpointer(x int, y int) {
	temp := x
	x = y
	y = temp

}




// now solve it using reference

// ChatGPT
// In Go, you don't explicitly use references like you would in languages such as C++. Go handles references implicitly when passing variables to functions. Here's how you can swap two numbers using this approach:

// In this code, swapByReference function takes two integers as arguments and returns them in reverse order, effectively swapping them. In the main function, we call swapByReference with the two numbers to swap their values and print the result. The concept of passing by reference in Go is implicit through the use of function parameters, unlike explicit references as seen in languages like C++.


func swapByReference(a, b int) (int, int) {
    return b, a
}











/*


User
func modifyValue(ptr *int) {
	   *ptr = 100 //Modify the value at the memory address pointed by 'ptr'
}

func main() {
    var num int = 42
    fmt.Println("Before:", num) // Before: 42
    modifyValue(&num)           // Pass the address of 'num' to the function
    fmt.Println("After:", num)  // After: 100
}

so in the above function basically this happens=>

var ptr *int=&num

right?
Answer=>

Yes, in the code snippet you provided, when you call the modifyValue(&num) function, you are passing the address of the variable num to the modifyValue function. This is done using the & operator, which obtains the memory address of a variable.

So, effectively, var ptr *int = &num happens internally when you pass &num as the argument to modifyValue. The ptr variable inside modifyValue then points to the memory address of num, allowing the function to modify the value stored at that address.

Here's the breakdown of what happens step by step:

var num int = 42: Declares and initializes an integer variable num with the value 42.
fmt.Println("Before:", num): Prints the value of num before calling the modifyValue function.
modifyValue(&num): Passes the address of num (&num) to the modifyValue function.
func modifyValue(ptr *int) { *ptr = 100 }: Receives the memory address of num as ptr and modifies the value at that address to 100.
fmt.Println("After:", num): Prints the value of num after the modifyValue function call, which is now 100 because the function modified the value at the memory address pointed to by ptr.
So, in essence, var ptr *int = &num doesn't need to be explicitly written in your code because Go automatically handles passing the address of num to the function when you use &num as an argument.

*/
