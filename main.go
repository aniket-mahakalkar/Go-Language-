package main

import "fmt"

// Variable in GO

func vrbls() {

	var age = 23
	teen_age := 19

	const gravity float64 = 9.8871 

	fmt.Println("The adult age is " , age ,"The teen age is", teen_age)
	fmt.Println("The gravity of earth is ",gravity)
}



// If Else in Go 
func return_age() {
	age :=  25

	if age >= 19 {

		fmt.Println("Your are an adult")

	}else if age >= 13 {

		fmt.Println("Your are in teenage")

	}else {

		fmt.Println("Your are child")

	}

	
	

}

// Loops in Go 
func loop() {

	fmt.Println("For Loop in Go ")
	for i := 0 ; i < 5 ; i++ {

		fmt.Println(i)
	}

	fmt.Println("While Loop in Go ")
	cnt := 0

	for cnt < 3 {

		fmt.Println(cnt)
		cnt++
	}
}

// Arrays and Slices in Go 

func return_array() {

	fmt.Println("Printing Array in Go")

	numbers := [] int {1,2,3,4,5}

	fmt.Println()
	fmt.Println("Printing using Println")
	fmt.Println(numbers)

	fmt.Println()
	fmt.Println("Printing using Loop")

	for ind, val := range numbers {

		fmt.Printf("Index %d Value : %d \n", ind, val)
	} 


	name := "GoLang"

	for _,char := range name {

		fmt.Printf("Char is : %c " , char)
	}
}


// Switch Case in Go

func switch_case() {

	fmt.Println("Switch Case Statment")
	grade := "B"

	switch grade {

	case "A" :
		fmt.Println("Excellent")

	case "B" :

		fmt.Println("Good")

	default :

		fmt.Println("Working Good")


}

}


func example_defer() {

	defer fmt.Println("Defer Execute Itself after all the neighbout function call Completed")

	fmt.Println("Aniket if ready")

	fmt.Println("Going forward to complete the match")
}


// Maps in Go 

func maps() {

	ages := map[string] int {"Alice": 30 , "Bob": 25}

	fmt.Println(ages["Alice"])

	ages["Charlie"] = 35

	delete(ages,"Bob")

	val , exits := ages["Charlie"]

	if exits {

		fmt.Println("Bob's ages:", val)


	}else {

		fmt.Println("Key 'Bob' not found.")


	}

	fmt.Println(len(ages))
}

func pointers() {

	age := 5

	var Pointer *int = &age
	fmt.Println(&age)
	fmt.Println(Pointer)
}

// Main Funtion 
func main() {

	fmt.Println("Welcome to Go Tutorial!")
	vrbls()
	return_age()
	loop()

	return_array()

	switch_case()

	example_defer()

	maps()

	fmt.Println("Practice with Pointers")

	pointers()

}