package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func greeter(name string) {
	fmt.Println("Welcom,", name)
}

func adder(one int, two int) int {
	return one + two
}

func proAdder(values ...int) int {
	sum := 0

	for _, val := range values {
		sum += val
	}

	return sum
}

func cStoI(elementStr string) (int64, error) {
	return strconv.ParseInt(strings.TrimSpace(elementStr), 10, 64)
}

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your good name :")
	input1, _ := reader.ReadString('\n')
	greeter(input1)

	fmt.Println("Enter first number :")
	input2, _ := reader.ReadString('\n')
	one, _ := cStoI(input2)

	fmt.Println("Enter second number :")
	input3, _ := reader.ReadString('\n')
	two, _ := cStoI(input3)

	addRes := adder(int(one), int(two))
	fmt.Printf("The addition of %d and %d is %d\n", one, two, addRes)

	fmt.Println("The result is :", proAdder(1, 2, 3, 4, 5, 6, 7, 87))
}
