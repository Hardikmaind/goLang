package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")
	content := "This needs to go in a file - LearnCodeOnline.in"

	file, err := os.Create("./mylcogofile.txt")

	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)

	length, err := io.WriteString(file, content)
	checkNilErr(err)
	fmt.Println("length is: ", length)
	defer file.Close()
	readFile("./mylcogofile.txt")
}

func readFile(filname string) {
	databyte, err := ioutil.ReadFile(filname)		//this is depricated, instead use os.ReadFile or os.Open
	filecontent, err := os.ReadFile(filname)		//used os.ReadFile 
	checkNilErr(err)

	fmt.Println("Text data inside the file is \n", string(databyte))
	fmt.Println("======================")
	fmt.Println("Text data inside the file is \n", string(filecontent))

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
