package main

import (
	"go-getting-started/syntax/basics/math"
	"fmt"
)

func main() {
	fmt.Print("My name is Dan\n")
	fmt.Printf("My name is %s. %T\n", "Dan", "Dan")

	var profession = "Biomedical Engineer"

	fmt.Println("I am a", profession)

	var exists = true

	if exists {
		println(exists)
	}

	var count = 3

	var ref = 2

	if count > 2 {
		println("equal to", ref)
	} else if count < 2 {
		println("equal to", ref)
	} else if count == 2 {
		println("equal to", ref)
	}

	var list1 = []string{"john", "steve", "rambo", "george"}

	println(list1, len(list1))

	for count := 0; count < len(list1); count++ {
		println(list1[count], count)
	}

	for count := range list1 {
		println(list1[count], count)
	}

	for count, value := range list1 {
		println(value, count)
	}

	var a = 20
	var b = 40

	var c int = math.Sum(a, b)

	fmt.Println("Sum = ", c)

	fmt.Println("Sequence: ",math.Fibonacci(1, 100))
	
	fmt.Println("Sum of Sequence= ",math.AddFibonacci(1, 100))
}
