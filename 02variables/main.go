package main

import "fmt"

const PI float32 = 3.14 // this is public

func main() {

	/*
		Variable Declaration :
			1. using var :
				syntax : var variable_name variable_type
			2. using implictly
				syntax : var variable_name = value
			3. without using var
				a) we can declare using this inside any method but globally
				   we cant declare that
				b) syntax : var variable_name := value
	*/

	// string type
	var username string = "dev.madhurendra"
	fmt.Println(username)
	fmt.Printf("Variable is of type : %T \n", username)

	// boolean type
	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type : %T \n", isLoggedIn)

	// integer -> unsigned int
	// uint8       the set of all unsigned  8-bit integers (0 to 255)
	// uint16      the set of all unsigned 16-bit integers (0 to 65535)
	// uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
	// uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)
	// integer -> signed int (with + and -)
	// uint8       the set of all unsigned  8-bit integers (-127 to 128)
	var smallInt uint8 = 255
	fmt.Println(smallInt)
	fmt.Printf("Variable is of type : %T \n", smallInt)

	// float -> float32
	// contains 5 value after decimal
	var smallFloat float32 = 255.443455445112334432232
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type : %T \n", smallFloat)

	// float -> float64
	// value contains only 0 - 255
	var bigFloat float64 = 255.443455445112334432232
	fmt.Println(bigFloat)
	fmt.Printf("Variable is of type : %T \n", bigFloat)

	// default values and some aliases
	var defaultVariable int
	fmt.Println("Default value of this variable : ", defaultVariable)

	// implicitly decide variable type
	var myAge = 18
	fmt.Println("myAge variable is declared implicitly : ", myAge)

	// without using a var
	email := "madhurendra.tiwari@zemosolabs.com"
	myBirthDate := 14
	mySalary := 42000.43
	isAdult := myAge >= 18
	fmt.Println("email,myAge,mySalary and isAdult variable declared without var and implicit : ", email, myBirthDate, mySalary, isAdult)

	fmt.Println("This is the value of PI which is constant : ", PI)

}
