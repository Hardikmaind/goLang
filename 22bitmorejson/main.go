package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`		//this is used to change the key name in the json data
	Price    int
	Platform string   `json:"website"`		//this is used to change the key name in the json data
	Password string   `json:"-"`			//this is used to not show the password key in the json data
	Tags     []string `json:"tags,omitempty"`			//this is used to not show the tags key if it is nil
}

func main() {
	fmt.Println("Welcome to JSON video")
	EncodeJson()
	fmt.Println("==========================================================================")
	DecodeJson()
}

func EncodeJson() {

	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "LearnCodeOnline.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "LearnCodeOnline.in", "bcd123", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 299, "LearnCodeOnline.in", "hit123", nil},
	}

	//package this data as JSON data

	// finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	// finalJson, err := json.Marshal(lcoCourses)			//this is just the normal json data
	finalJson, err := json.MarshalIndent(lcoCourses, "", "  ")				//here the third argument is the spacing in the indecntation 
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finalJson)

}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJS Bootcamp",
		"Price": 299,
		"website": "LearnCodeOnline.in",
		"tags": ["web-dev","js"]
	}
	`)

	var lcoCourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)			//func json.Unmarshal(data []byte, v any) error
		// Unmarshal parses the JSON-encoded data and stores the result in the value pointed to by v. If v is nil or not a pointer, Unmarshal returns an [InvalidUnmarshalError]
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON WAS NOT VALID")
	}

	// some cases where you just want to add data to key value

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)			//in order to print the map data we need to use the %#v. here we are print the interface data(map[string]interface{})

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v  and value is %v and Type is: %T\n", k, v, v)
	}

fmt.Println("==========================================================================")

	var hm=make(map[string]string)	
	hm["name"]="saurabh"
	hm["age"]="22"
	hm["city"]="delhi"
	hm["country"]="india"
	fmt.Printf("this is the mapped data=>>>>>> %#v \n  ",hm)		//This line prints the map hm2 using %#v, which is a verb specifically designed to provide a Go-syntax representation of the value. For a map with string keys and integer values (map[string]int), %#v prints the map literal in a format that can be directly used as Go code to recreate the map.

	fmt.Printf("this is the mapped data=>>>>>> %v \n  ",hm)	//This line prints the map hm using %v, which is a general-purpose verb that prints the default format for each type. For a map with string keys and string values (map[string]string), %v prints the key-value pairs in a readable format, similar to how you would define a map literal in Go code.


     fmt.Println("=============================this is the map of type string and key and =============================================")





	hm2 := map[string]int{
        "apple":  5,
        "banana": 3,
        "cherry": 8,
    }

    // Print the map using %#v
    fmt.Printf("Map: %#v\n", hm2)
}
