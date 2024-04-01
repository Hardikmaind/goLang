package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang; No super or parent

	hitesh := User{"Hitesh", "hitesh@go.dev", true, 16}
	fmt.Println(hitesh)			//this is without %+v
	fmt.Printf("hitesh details are: %+v\n", hitesh)		//this is with %+v
	fmt.Printf("Name is %v and email is %v.", hitesh.Name, hitesh.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
