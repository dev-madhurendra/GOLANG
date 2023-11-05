package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	start := time.Now().Second()
	fmt.Println("Enter your username : ")

	username, _ := reader.ReadString('\n')

	end := time.Now().Second()
	timeTaken := end - start
	fmt.Println("Hey,"+username+"you took ", timeTaken, "second to enter your username !")

	t := time.Now().Format(time.UnixDate)
	fmt.Println("Time :", t)

	formattedTime := time.Now().Format(time.DateOnly)
	fmt.Println("Custom formatted time : ", formattedTime)

	formattedTimeDateDay := time.Now().Format("2006-01-02 Monday")
	fmt.Println("Custom formatted time date day : ", formattedTimeDateDay)

	formattedTimeDateDayTime := time.Now().Format("2006-01-02 15:01:01 Monday")
	fmt.Println("Custom formatted time date day time : ", formattedTimeDateDayTime)
}
