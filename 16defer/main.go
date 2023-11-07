package main

import "fmt"

func myDefer() {
	for i := 0; i < 6; i++ {
		defer fmt.Print(i, " ")
	}
}

func main() {

	// the output will be Hello,World !
	defer fmt.Println("World !")
	fmt.Print("Hello,\n")

	// the output will be 1 3 4 2
	fmt.Println("1")
	defer fmt.Println("2")
	fmt.Println("3")
	defer fmt.Println("4")

	myDefer()

	defer myDefer()
}
