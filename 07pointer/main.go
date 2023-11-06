package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your marks = ")
	input, _ := reader.ReadString('\n')

	marks, _ := strconv.ParseInt(strings.TrimSpace(input), 10, 64)

	var ptr = &marks

	fmt.Println("Value of actual pointer is ", ptr)
	fmt.Println("Value of actual pointer is ", *ptr)

	*ptr = *ptr * 2
	fmt.Println("New value is ", *ptr)
}
