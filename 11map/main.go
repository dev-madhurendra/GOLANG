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

	m_a_p := make(map[string]int)

	m_a_p["Saiprabhu"] = 18000
	m_a_p["Anna Hariprasad"] = 42000
	m_a_p["Madhurendra"] = 18000

	// Map value before change
	fmt.Println("This is the map (before) : ", m_a_p)

	// After performance consideration
	m_a_p["Madhurendra"] = 42000

	// Map value after change
	fmt.Println("This is the map (after) : ", m_a_p)

	// Iterating map using for range loop
	for key, value := range m_a_p {
		fmt.Println(key, " =", value)
	}

	/*
		Use Case : find the uniqe element from an array
	*/

	m_a_p_numbers := make(map[int]int)

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter the length of an array = ")
	input, _ := reader.ReadString('\n')
	length, _ := cStoI(input)
	numbers := make([]int64, length)

	for i := int64(0); i < length; i++ {
		fmt.Printf("Enter element %d: ", i)
		elementStr, _ := reader.ReadString('\n')
		element, err := cStoI(elementStr)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		numbers[i] = element
		m_a_p_numbers[int(element)]++
	}

	for key, value := range m_a_p_numbers {
		fmt.Println("Key ", key, " = ", value)
	}

}
