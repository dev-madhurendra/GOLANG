package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var username string
	welcome := "Welcom ," + username
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Hey, enter you username : ")

	// comma ok || err ok
	inputString, _ := reader.ReadString('\n')
	fmt.Println("Welcom , ", inputString)
}
