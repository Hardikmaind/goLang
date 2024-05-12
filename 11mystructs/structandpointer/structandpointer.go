package main

import "fmt"

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
	fmt.Println(hardik)
}
