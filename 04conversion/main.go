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

	fmt.Print("Enter your marks: ")
	inputMarks, _ := reader.ReadString('\n')

	marks, _ := strconv.ParseInt(strings.TrimSpace(inputMarks), 10, 16)

	totalPercentage := (float32)(marks*100) / 600
	fmt.Println("Total Percentage:", totalPercentage, "%")
}
