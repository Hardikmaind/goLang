package main

import "fmt"

type User struct {
    Name   string
    Email  string
    Status bool
    Age    int
}

// Method with a value receiver
func (u User) printDetails() {
    fmt.Printf("Name: %s, Email: %s, Status: %t, Age: %d\n", u.Name, u.Email, u.Status, u.Age)
}

// Method with a pointer receiver
func (u *User) updateEmail(newEmail string) {
    u.Email = newEmail
}

// Method with a pointer receiver
func (u *User) updateUserDetails(name, email string, status bool, age int) {
    u.Name = name
    u.Email = email
    u.Status = status
    u.Age = age
}


func main() {
    // Creating an instance of User
    temp := User{Name: "Alice", Email: "alice@example.com", Status: true, Age: 30}

    // Calling a method with a value receiver
    temp.printDetails()

    // Calling a method with a pointer receiver
    temp.updateEmail("newemail@example.com")
	// No, temp is not a pointer to User in this code. It's an instance of the User struct. In Go, when you define a method with a pointer receiver, Go automatically converts the variable to a pointer when you call the method on a value. This conversion happens implicitly behind the scenes.

	// So, when you call temp.updateEmail("newemail@example.com"), Go automatically converts temp to a pointer (&temp) and passes it to the updateEmail method, allowing the method to update the Email field of the original User object referred to by temp.
	



    // Since updateEmail has a pointer receiver, Go automatically converts the variable to a pointer
    // and updates the Email field of the original variable.
    temp.printDetails()









	 // Create a User object
	 temp2 := User{}

	 // Create a pointer to the User object
	 userPtr := &temp2
 
	 // Update user details using the pointer
	 userPtr.updateUserDetails("Hardik", "maindhardik@gmail.com", true, 23)
	//  Yes, userPtr is a pointer to the User object temp2 in this code. By using the & operator, you're taking the address of temp2 and assigning it to userPtr, making userPtr a pointer to temp2. This allows you to modify temp2 through userPtr, as demonstrated by the updateUserDetails method call.



	
 
	 // Print updated user details
	 fmt.Printf("Name: %s, Email: %s, Status: %t, Age: %d\n", temp2.Name, temp2.Email, temp2.Status, temp2.Age)
}
