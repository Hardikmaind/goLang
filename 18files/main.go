package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "Hello my name is hardik maind."

	file, err := os.Create("18files/mylcogofile.txt")

	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)

	length, err := io.WriteString(file, content) //io.WriteString is used to write the string to the file. it returns the length of the string written and error if any
	checkNilErr(err)
	fmt.Println("length is: ", length)
	defer file.Close() //defer is used to close the file after the execution of the function. this will be used at the end of the function
	readFile("18files/mylcogofile.txt")
}

func readFile(filname string) {
	databyte, err := ioutil.ReadFile(filname) //this is depricated, instead use os.ReadFile or os.Open
	filecontent, err := os.ReadFile(filname)  //used os.ReadFile
	//whenever we are reading the file, it is not being read into the string format and that is the biggest gotcha moment in we always need to keep an eye specially when we are reading data from the internet as well, that is always read into the bytes formatccc
	checkNilErr(err)

	fmt.Println("Text data inside the file is \n", string(databyte))
	fmt.Println("======================")
	fmt.Println("Text data inside the file is =>>. it is in the byte format we have type casted it\n", string(filecontent))
	fmt.Print("this is the byte format ....", filecontent)

}

func checkNilErr(err error) {
	if err != nil {
		panic(err) //panic is used to stop the execution of the program and show the error
	}
}
