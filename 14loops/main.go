package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for _, day := range days {
	// 	fmt.Printf("index is  and value is %v\n", day)
	// }

	for index, day := range days { //this is like the for each loop in golang.		first value is the index and second value is the value of the array
		fmt.Println("Index is ", index, " and value is ", day)
	}

	rougueValue := 1

	for rougueValue < 10 {

		if rougueValue == 3 {
			rougueValue++
			continue
		}

		if rougueValue == 2 {
			goto lco
		}

		if rougueValue == 5 {
			rougueValue++
			continue
		}

		fmt.Println("Value is: ", rougueValue)
		rougueValue++
	}

lco:
	fmt.Println("Jumping at LearnCodeonline.in")

	// time package
	// this is the code i have written to check the time taken by the loop to execute

	start := time.Now()
	fmt.Println("Start time is ", start)
	i := 0
	sum := 0
	for i < 10000 {
		sum += i
		i++
	}
	end := time.Now()
	fmt.Println("End time is ", end)
	fmt.Println("start time is ", start)
	totaltime := end.Sub(start)
	fmt.Println("Total time taken is ", totaltime)

}
