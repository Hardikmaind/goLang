package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("List of all languages: ", languages)
	fmt.Println("JS shorts for: ", languages["JS"])

	delete(languages, "RB")
	fmt.Println("List of all languages: ", languages)

	// loops are interesting in golang

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v\n",key, value)
	}
	for key := range languages {
		fmt.Printf("For key %v, value is v\n",key )
	}

	// above and below same thing
	// for key ,_:= range languages {
	// 	fmt.Printf("For key %v, value is v\n",key )
	// }



	for  value := range languages {
		fmt.Printf("For key v, value is %v\n", value)
	}

	// above and below same thing
	// for _, value := range languages {
	// 	fmt.Printf("For key v, value is %v\n", value)
	// }

}
