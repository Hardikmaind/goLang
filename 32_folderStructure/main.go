package main

// create a module first with the cmds go mod init github.com/Hardikmaind/<anything>
//And then we are ready to this any of importing.This is how we can import the packages from different folders in go
import ("fmt"
		"github.com/Hardikmaind/go_folder_structure/f2"
		"github.com/Hardikmaind/go_folder_structure/f1/f3"
		"github.com/Hardikmaind/go_folder_structure/f1"
	)
func main(){
	fmt.Println("Hello World")
	f2.Demo()
	f3.Demo()
	f1.Demo()
}