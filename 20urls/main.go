package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"

func main() {
	fmt.Println("Welcome to handling URLs in golang")
	fmt.Println(myurl)

	//parsing
	result, _ := url.Parse(myurl) //we need to parse the url to get the parts of the url

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	// this beliw is the ouput of the above code

	// 	https
	// lco.dev:3000
	// /learn
	// 3000

	qparams := result.Query()
	fmt.Printf("The type of query params are: %T\n", qparams)

	fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("Param is: ", val)
	}




	// this is the part where you have all the chunks of the url and we have to contruct the url from it
	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}
	fmt.Print("this is the url=>  ",partsOfUrl,"\n")
	fmt.Printf("type of partsOfUrl is: %T",partsOfUrl)
	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)

}
