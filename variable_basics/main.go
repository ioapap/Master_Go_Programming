/////////////////////////////////
// Variables and Declarations
// Go Playground: https://play.golang.org/p/PKdAxUp8mNT
/////////////////////////////////

package main

import "fmt"

func main() {
	// int age = 10; //c++ way
	var age int = 21
	fmt.Println("Age:", age)

	var name = "Dan"
	// commenting the line below just to produce an error and understand how the "blank identifier" works
	// fmt.Println("Your name is:", name)
	_ = name
	//☝ this is called : "blank identifier"

	s := "Learning golang!"
	fmt.Println(s)

	// cannot use ":=" for already defined variables
	// the line below is an error, "no new variables on left side of :=go"
	// name := "Andrei"

	// if we want to change the value of the name variable we use only the "=" sign but if we want to create a new variable for example:
	// name1 := John //there is no error here
	// _ = name1     //blank identifier

	car, cost := "Maserati", 50000
	fmt.Println(car, cost)
	// If we try to re-declare these variables in the same way we will get an error message
	// Example: car, cost := "Maserati", 60000 -> "no new variables on left side of :=go"
	// We can use re-declaration with short declaration syntax ONLY when at least one variable on the left hand side of := is new
	// Example: car, year := "Maserati", 2018 -> this is not an error, this is an existing variable with a new one, at least one is new

	var opened = false
	opened, file := true, "a.txt"

	_, _ = opened, file

	// Multiple declaration which is good for readability
	// Another way to declare multiple variables at once using normal declaration or the "var" keyword

	var (
		salary    float64
		firstName string
		gender    bool
	)
	fmt.Println(salary, firstName, gender)

	// Another concise way to declare multiple variables that have the same type

	var a, b, c int
	fmt.Println(a, b, c)

	// When do we use short declaration and when do we use normal declaration?
	// We use short declaration so := when we know the initial value and normal declaration the var keyword otherwise.
	// We also use normal declaration for multiple declaration for a better readability.

	// Multiple assignment

	var i, j int
	i, j = 5, 8

	// j, i = i, j // swapping variables

	_, _ = i, j

	// It's possible to use expressions in short declarations

	sum := 5 + 2.3
	fmt.Println(sum)

}
