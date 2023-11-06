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

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your age : ")
	input1, _ := reader.ReadString('\n')
	myAge, _ := cStoI(input1)

	if myAge < 18 {
		fmt.Println("You can not vote !")
	} else {
		fmt.Println("You can vote !")
	}

	fmt.Println("Enter any value :")
	input2, _ := reader.ReadString('\n')
	value, _ := cStoI(input2)

	if value < 0 {
		fmt.Println("Value is negative !")
	} else if value == 0 {
		fmt.Println("Value is zero !")
	} else {
		fmt.Println("Value is positive !")
	}

	fmt.Println("Enter the day : ")
	input3, _ := reader.ReadString('\n')
	day, _ := cStoI(input3)

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Weekday")
	default:
		fmt.Println("Weekday")
	}

	// golang does not have ternary operator
	// isAdult := (myAge >= 18) ? true : false
}
