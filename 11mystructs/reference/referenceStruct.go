package reference

import "fmt"

type Dummy struct{
	city string
	country string
	state string
}

func Demo() {
	fmt.Println("")
	fmt.Println("This is a Demo function")
	fmt.Println("")
	fmt.Println("hello from mystructs/reference/referenceStruct.go")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
}

func Test() {
	fmt.Println("")
	fmt.Println("This is a Test function")
	fmt.Println("")
	fmt.Println("Hello from from mystructs/reference/referenceStruct.go")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
	fmt.Println("")
}

func AddCity( city string, country string, state string) *Dummy {
	return &Dummy{city: city, country: country, state: state}
}
