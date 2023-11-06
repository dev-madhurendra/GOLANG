package main

import "fmt"

func main() {

	var fruits = []string{"Apple", "Mangoe", "Orange"}
	fmt.Printf("Type of fruits is %T\n", fruits)

	fruits = append(fruits, "Papaya", "Guvava")
	fmt.Println("The list of fruit : ", fruits)

	newFruits := append(fruits[3:5], "Grapes")
	fmt.Println("The list of new fruits : ", newFruits)

	// slice
	/*
		Ques : Slice an array from 2 to 4 position
		input : [1,2,3,4,5]
		output : [3,4]

	*/

	numbers := []int{1, 2, 3, 4, 5}
	newNumbers := numbers[2:4] // [3 4]
	fmt.Println("Slice of numbers list is : ", newNumbers)

}
