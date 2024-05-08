package main

import "fmt"

//CONCEPT OF POINTER IS DIFFERENT IN GO AS COMPARED TO C++.!!!!!!!!!!
//CONCEPT OF POINTER IS DIFFERENT IN GO AS COMPARED TO C++.!!!!!!!!!!
//CONCEPT OF POINTER IS DIFFERENT IN GO AS COMPARED TO C++.!!!!!!!!!!
//CONCEPT OF POINTER IS DIFFERENT IN GO AS COMPARED TO C++.!!!!!!!!!!
//CONCEPT OF POINTER IS DIFFERENT IN GO AS COMPARED TO C++.!!!!!!!!!!
//CONCEPT OF POINTER IS DIFFERENT IN GO AS COMPARED TO C++.!!!!!!!!!!

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang; No super or parent

	hitesh := User{"Hitesh", "hitesh@go.dev", true, 16}
	fmt.Println(hitesh)                             //this is without %+v
	fmt.Printf("hitesh details are: %+v\n", hitesh) //this is with %+v
	fmt.Printf("Name is %v and email is %v.", hitesh.Name, hitesh.Email)

	fmt.Println("")
	user := User2{Name: "Hardik", Email: "maindhardik@gmail.com", address: "Mumbai", gender: "male"}
	fmt.Println("User2 is: ", user)

	fmt.Println("User2 gender is: ", user.getGender())
	fmt.Println("User2 adderess is: ", user.getAddress())
	user.updateAddress("Pune")		//here go automatically dereferences the pointer which in this case is user and then accesses the address field of the User2 struct
	fmt.Println("User2 this is the updated adderess is: ", user.getAddress())

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

type User2 struct {
	Name    string
	Email   string
	address string
	gender  string
}

func (u User2) getGender() string {			//now here the whole struct is passed as a value means that the struct is copied and passed to the function
	return u.Name
}

func (u *User2) getAddress() string {		//now here the whole struct is passed as a pointer means that the address of the struct is passed to the function
	return u.address
}


// In Go, when you use a pointer receiver in a method, you don't explicitly dereference the pointer within the method body because Go automatically handles that for you. When you call a method on a pointer receiver, Go automatically dereferences the pointer and accesses the underlying struct.

// In your example, func (u *User2) updateAddress(address string), u is a pointer to User2. When you call user.updateAddress("New Address"), Go understands that user is a pointer and automatically dereferences it to access the User2 struct's fields and methods.

// So, when you write u.address = address inside the updateAddress method, Go automatically dereferences u and accesses the address field of the User2 struct. This behavior simplifies the code and makes it more readable without the need for explicit dereferencing.

func (u *User2)updateAddress(address string){
	u.address = address		//here no need to use *u.address because the struct is passed as a pointer and the address is accessed directly. go automatically dereferences the pointer and accesses the underlying struct. this is not cpp see cpp code below to swap numbers
}


// unlike in cpp
// #include <iostream>

// void swapNumbers(int* ptr1, int* ptr2) {
//     int temp = *ptr1;
//     *ptr1 = *ptr2;
//     *ptr2 = temp;
// }

// int main() {
//     int num1 = 10, num2 = 20;
//     int* ptr1 = &num1;
//     int* ptr2 = &num2;

//     std::cout << "Before swapping: num1 = " << *ptr1 << ", num2 = " << *ptr2 << std::endl;
//     swapNumbers(ptr1, ptr2);
//     std::cout << "After swapping: num1 = " << *ptr1 << ", num2 = " << *ptr2 << std::endl;

//     return 0;
// }
