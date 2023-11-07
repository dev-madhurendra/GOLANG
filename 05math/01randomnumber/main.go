package main

import (
	"fmt"
	"math/rand"
)

func main() {
	randomNumber := rand.Intn(6) + 1
	fmt.Println("Dice number is : ", randomNumber)
}
