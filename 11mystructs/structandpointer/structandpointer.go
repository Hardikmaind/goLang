package main

import (
	"fmt"
	"mystructs/reference"
)

type Student struct {
	Name   string
	RollNo int
	// Add more fields as needed
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func AddStudent(students *[]Student, name string, rollNo int) {
	// Create a new student
	student := Student{
		Name:   name,
		RollNo: rollNo,
		// Set other fields as needed
	}

	// Append the new student to the slice
	*students = append(*students, student)
}
func (h *User) addUserDetals(Name string, Email string, status bool, Age int) {
	h.Email = Email
	h.Name = Name
	h.Status = status
	h.Age = Age
}
func main() {
	var students []Student

	// Add a new student
	AddStudent(&students, "Alice", 101) //this is the function

	// Print the students
	for i, student := range students {
		fmt.Printf("Student %d: Name=%s, RollNo=%d\n", i+1, student.Name, student.RollNo)
	}

	// Create a new user
	hardik := User{}
	fmt.Println(hardik) //create a hardik object of User struct but no values are added into it
	// her i have used the pointer receiver method to add values to the hardik object.
	// this is the method of the User struct and the receiver is a pointer to the User struct.
	hardik.addUserDetals("Hardik", "maindhardik@gmail.com,", true, 23) //add values to the hardik object .
	//  in the main function, hardik is not a pointer to a struct (*User), it is an instance of the User struct. The method addUserDetails has a receiver of type *User, which means it can be called on a pointer to a User struct. When you call hardik.addUserDetails(...), Go automatically takes the address of hardik (&hardik) and passes it as the receiver to the method, allowing the method to modify the fields of the hardik object directly. This is why you can call a pointer receiver method on a non-pointer (User) type without explicitly taking its address.







	fmt.Println(hardik)

	reference.Test()
	fmt.Println("this is M-I ===>>")
	var dummy *reference.Dummy //sqiggle because  it wants use to write like this=> 	dummy := reference.AddCity("Pune", "India", "MH")  ;    here the dummy will be infered as a pointer to the Dummy struct

	dummy = reference.AddCity("Pune", "India", "MH")
	fmt.Println(&dummy) //this will print the address of the dummy pointer
	fmt.Println(dummy)  //this will print the value of the dummy pointer
	fmt.Println(*dummy, "\n\n")

	fmt.Println("this is M-II ====>>>")
	temp := new(reference.Dummy) //The new function allocates memory for a new zero-initialized value of the specified type and returns a pointer to it. So, dummy will be a pointer to a new Dummy struct.
	temp = reference.AddCity("Pune", "India", "MH")
	fmt.Println(temp, "\n\n")

	// new function in Go is used to allocate memory for a new zero-initialized value of a specified type. It returns a pointer to the newly allocated memory. Since int is a basic type in Go and not a struct type, you cannot use new to create a new int variable directly.

	// var i int             // Declares an int variable with a zero value (0)
	j := 42        // Declares and initializes an int variable with the value 42
	k := new(int)  // Declares a pointer to an int and initializes it with nil
	fmt.Println(k) // Prints the memory location of the pointer k
	*k = j         // Assigns a value of 100 to the memory location pointed by k
	fmt.Println("Value of k:", *k)

}
