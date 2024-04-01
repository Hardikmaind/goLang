package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang; No super or parent

	hitesh := User{"Hitesh", "hitesh@go.dev", true, 16}
	fmt.Println(hitesh)
	fmt.Printf("hitesh details are: %+v\n", hitesh)
	fmt.Printf("Name is %v and email is %v.\n", hitesh.Name, hitesh.Email)
	hitesh.GetStatus()
	hitesh.NewMail()
	fmt.Printf("Name is %v and email is %v.\n", hitesh.Name, hitesh.Email)
	fmt.Println("================setter created by me are below=================\n\n\n")



	name:="HARDIK MAIND"
	hitesh.AddName(name)
	// now here the name will not be changed because the name is not changed in the original struct. here we are doing pass by value
	fmt.Println("Name of this user is: ", hitesh.Name)
	fmt.Println("======AddName1 above===========")

	// here the name will be changed because the name is changed in the original struct . becauese pointer of the struct is passed. here we are doing pass by reference

	hitesh.AddName2(name)
	fmt.Println("Name of this user is: ", hitesh.Name)
	fmt.Println("======AddName2 above===========")


}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is: ", u.Email)
}

// this is created by me .
// this settter method will not change the name of the user because the name is not changed in the original struct.why? because we are passing the value of the struct
func (u User)AddName(data string){
	u.Name=data
	fmt.Println("Name of this user is: ", u.Name)
}

// this will change the name of the user because the name is changed in the original struct. why ? because we are passing the pointer to the struct
func (u *User)AddName2(data string){
	u.Name=data
	fmt.Println("Name of this user is: ", u.Name)
}
