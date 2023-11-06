package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func cStoI(elementStr string) (int64, error) {
	return strconv.ParseInt(strings.TrimSpace(elementStr), 10, 64)
}

func main() {

	// fixed length
	var fruits [3]string
	fruits[0] = "Mangoes"
	fruits[1] = "Apple"
	fruits[2] = "Banana"

	fmt.Println("List of fruits : ", fruits)

	var vegetables = [3]string{"Potato", "Tomato", "Lady finger"}
	fmt.Println("List of vegetables : ", vegetables)

	// Create a bufio.Reader to read from standard input (os.Stdin)
	reader := bufio.NewReader(os.Stdin)

	// Ask the user to specify the length of the slice
	fmt.Print("Enter the length of the slice: ")
	lengthStr, _ := reader.ReadString('\n')

	length, err := cStoI(lengthStr)

	// Create a slice to store the user's input values
	// var numbers []int64
	numbers := make([]int64, length)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Read and store user input data
	for i := int64(0); i < length; i++ {
		fmt.Printf("Enter element %d: ", i)
		elementStr, _ := reader.ReadString('\n')
		element, err := cStoI(elementStr)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		numbers[i] = element
		// numbers = append(numbers, element)
	}

	// Print the resulting slice
	fmt.Println("Your slice:", numbers)
}
