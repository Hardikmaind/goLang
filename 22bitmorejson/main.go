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
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v  and value is %v and Type is: %T\n", k, v, v)
	}

}
