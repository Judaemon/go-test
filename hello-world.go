package main

import "fmt"

func main() {
    fmt.Println("hello world")

	fmt.Println("go" + "lang")

    fmt.Println("1+1 =", 1+1)
    fmt.Println("7.0/3.0 =", 7.0/3.0)

    fmt.Println(true && false)
    fmt.Println(true || false)
    fmt.Println(!true)

	var a = "initial"
    fmt.Println(a)

	// multiple variables can be declared at once
    var b, c int = 1, 2
    fmt.Println(b, c)

	// Go will infer the type of initialized variables
    var isHappy = true
    fmt.Println(isHappy)

	// Variables declared without a corresponding initialization are zero-valued.
    var score int
    fmt.Println(score)

	// The := syntax is shorthand for declaring and initializing a variable
    var fruit = string("apple")
    fmt.Println(fruit)
}