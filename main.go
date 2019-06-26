package main

import "fmt"

func main() {


	//================================ 1

	//fmt.Println("Hello, World")

	//================================ 2

	// Declaring Variables
	// var myStr string = "Hello"
	// var myInt int = 100
	// var myFloat float64 = 45.12
	//fmt.Println(myStr, myInt, myFloat)


	// Multiple Declarations
	// var (
	// 	employeeId int = 5
	// 	firstName, lastName string = "Satoshi", "Nakamoto"
	// )
	//fmt.Println(employeeId, firstName, lastName)

	
	// Short variable declaration syntax
	// name := "Rajeev Singh"
	// age, salary, isProgrammer := 35, 50000.0, true

	//fmt.Println(name, age, salary, isProgrammer)


	//================================ 3
	// var name = "Rajeev Singh" // Type declaration is optional here.
	// fmt.Printf("Variable 'name' is of type %T\n", name)

	// //================================

	// // Multiple variable declarations with inferred types
	// var firstName, lastName, age, salary = "John", "Maxwell", 28, 50000.0

	// fmt.Printf("firstName: %T, lastName: %T, age: %T, salary: %T\n",
	// 	firstName, lastName, age, salary)

	//================================= 4

	var (
		firstName, lastName string
		age                 int
		salary              float64
		isConfirmed         bool
	)

	fmt.Printf("firstName: %s, lastName: %s, age: %d, salary: %f, isConfirmed: %t\n",
		firstName, lastName, age, salary, isConfirmed)

	//================================== 5
	var myInt8 int8 = 97

	/*
	  When you don't declare any type explicitly, the type inferred is `int`
	  (The default type for integers)
	*/
	var myInt = 1200

	var myUint uint = 500

	var myHexNumber = 0xFF  // Use prefix '0x' or '0X' for declaring hexadecimal numbers
	var myOctalNumber = 034 // Use prefix '0' for declaring octal numbers

	var myFloat32 float32 = 4.5
	var myFloat = 9.12 // // Type inferred as `float64` (the default type for floating-point numbers)

	fmt.Printf("%d, %d, %d, %#x, %#o %f %f\n", myInt8, myInt, myUint, myHexNumber, myOctalNumber, myFloat32, myFloat)

}