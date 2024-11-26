package main //!this is the main package..And in f4 folder also we have main package....so if the demo2.go file was in the root folder then we will see the error because we have two main packages(also two main functions) in the same folder. so if they are kept in different folder then it will not give any error.

// create a module first with the cmds go mod init github.com/Hardikmaind/<anything>
//And then we are ready to this any of importing.This is how we can import the packages from different folders in go
import (
	"fmt"

	"github.com/Hardikmaind/go_folder_structure/f1"
	"github.com/Hardikmaind/go_folder_structure/f1/f3"
	"github.com/Hardikmaind/go_folder_structure/f2"
)

func main() {
	fmt.Println("Hello World")
	f2.Demo()
	f3.Demo()
	f1.Demo()
}
