package main

import "fmt"

func main() {
	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Satruday"}

	fmt.Println(days)

	for d := 0; d < len(days); d++ {
		fmt.Println("Day ", days[d])
	}

	for i := range days {
		fmt.Println(days[i])
	}

}
