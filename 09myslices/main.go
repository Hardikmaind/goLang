package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to video on slices")

	var fruitList = []string{"Apple", "Tomato", "Peach"}
	fmt.Printf("Type of fruitlist is %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)
	fmt.Println("=====================")
	fruitList = append(fruitList[0:3]) //here 0 is inclusive and 3 is exclusive
	fmt.Println(fruitList)

	fruitList = append(fruitList[0:]) //here 1 is inclusive and 2 is exclusive
	fmt.Println(fruitList)            //now since we have already sliced the slice so it will print the same slice and elements in it

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	//highScores[4] = 777
	fmt.Println("before appending ", len(highScores))

	highScores = append(highScores, 555, 666, 321)

	fmt.Println(highScores)
	fmt.Println("this is the length of the slice highscores ",len(highScores))

	//fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	//fmt.Println(highScores)

	//how to remove a value from slices based on index

	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
