package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our Pizza:")

	// comma ok || err err

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating, ", input)
	fmt.Printf("Type of this rating is %T", input)
	ok()
	for i := 0; i < 5; i++ {
		fmt.Println("Value of i is", i)

	}
	fmt.Println("hgere i hace called te function to print the table of 25")
	fmt.Println("Enter a number to print table:")
	num := reader.Buffered()
	print2Table(num)


	// Using fmt.Scanln(). this is use take string input and store in the variable. if any space is there then it will take only first word
	name := takeInputScanln()
	fmt.Println("Hello,", name)

	// Using fmt.Scanf(). this is use to take integer input and store in the variable
	age := takeInputScanf()
	fmt.Println("You are", age, "years old")

	// Using bufio.NewReader() and reader.ReadString(). this is use to take string input and store in the variable. if any space is there then it will take whole sentence
	sentence := takeInputReadString()
	fmt.Println("You entered:", sentence)

	// Using bufio.NewScanner(). this is use to take multiple lines of input and store in a slice
	lines := takeInputNewScanner()
	fmt.Println("You entered multiple lines:")
	for _, line := range lines {
		fmt.Println(line)
	}
}

func ok() {
	fmt.Println("")
	temp := "this is the OK function called here"
	fmt.Println("function called=>", temp)
}
func print2Table(num int) {
	fmt.Println("Printing table for", num)
	for i := 1; i <= 10; i++ {
		fmt.Println(num, "X", i, "=", num*i)
	}
}

// Function to take input using fmt.Scanln()
func takeInputScanln() string {
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)
	return name
}

// Function to take input using fmt.Scanf()
func takeInputScanf() int {
	var age int
	fmt.Print("Enter your age: ")
	fmt.Scanf("%d", &age)
	return age
}

// Function to take input using bufio.NewReader() and reader.ReadString()
func takeInputReadString() string {
	fmt.Print("Enter a sentence: ")
	reader := bufio.NewReader(os.Stdin)
	sentence, _ := reader.ReadString('\n')
	return sentence
}

// Function to take input using bufio.NewScanner()
func takeInputNewScanner() []string {
	var lines []string
	fmt.Println("This can be achieved by pressing Ctrl+D on Unix-based systems or Ctrl+Z on Windows systems, followed by pressing Enter. This key combination signals the end of input (EOF - End of File) to the scanner, causing it to stop reading input.")
	fmt.Println("Enter multiple lines (press Ctrl+Z to finish):")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	return lines
}
