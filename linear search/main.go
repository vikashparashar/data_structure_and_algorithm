package main

import (
	"fmt"
	"os"
	"strconv"
)

// lets say that we have to guss a number from range of 1 to 100
// and guss number will we provided with args.

// onec we find that guss number we will print that on concole
// and also print the steps taked by our program to find that guss number

// we will do this in a way of linear search algorithm

// as a linear search we will start to find that number from 1 to upto 100 one by one
// this algorithm will take 100 seps to find that number if it is 100
// take 33 steps to find that number if the gussed number is 33

// this algorithm to find that number is called linear search
func main() {
	var count = 1
	gussnumstr := os.Args[1]
	gussnum, _ := strconv.Atoi(gussnumstr)
	//
	if gussnum > 100 {
		fmt.Println("Please select number in range btw 1 to 100")
		os.Exit(1)
	}
	for i := 1; i <= 100; i++ {
		if i == gussnum {
			fmt.Println("guss number is : ", i)
			fmt.Println("Setps Taken To Find The Guss Number Are : ", count)
			os.Exit(1)
		} else {
			fmt.Println(count, "Searching Again ...")
		}
		count++
	}
}
