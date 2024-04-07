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

	fmt.Println("==================================================================")
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port()) 
	fmt.Println(result.RawQuery)			//same as in javascript it is knows as params or parameters	
	fmt.Println("==================================================================")
	// this beliw is the ouput of the above code

	// 	https
	// lco.dev:3000
	// /learn
	// 3000

	qparams := result.Query()		//.Quert() just stores the query parameters in the form of key value pair. from this we can get the values of the query parameters
	fmt.Printf("The type of query params are: %T\n", qparams)

	fmt.Println(qparams["coursename"])
 
	for _, val := range qparams {
		fmt.Println("value or  Param is: ", val)
	}

	for key, _ := range qparams {
		fmt.Println("key  is: ", key)
	}




	// this is the part where you have all the chunks of the url and we have to contruct the url from it
	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}
	fmt.Print("this is the url=>  ",partsOfUrl,"\n")
	fmt.Println("type of partsOfUrl is: %T",partsOfUrl)
	anotherURL := partsOfUrl.String()		// we can also it in this way ==>> anotherURL := String(partsOfUrl 	)
	fmt.Println("this is  url but in string ",anotherURL)

}
