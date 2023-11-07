package main

import "fmt"

func main() {
	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Satruday", "Sunday"}

	fmt.Println(days)

	for d := 0; d < len(days); d++ {
		fmt.Println("Day ", days[d])
	}

	for i := range days {
		if days[i] == "Saturday" || days[i] == "Sunday" {
			fmt.Println("Weekdays")
			continue
		} else {
			fmt.Println("Working day : ", days[i])
		}
	}

	rogueValue := 1

	for rogueValue < 10 {

		if rogueValue == 7 {
			rogueValue++
			goto goto_statement
			// continue
		} else if rogueValue == 9 {
			break
		} else {
			fmt.Println("Value is : ", rogueValue)
		}
		rogueValue++
	}
goto_statement:
	fmt.Println("Jumping to the line number 38")
}
