package main

import "fmt"

// Var makes variables accessable packagewise.
var a string
var b int
var c float64
var d bool

// Shorthand notations ":=" can ONLY be used inside func
func main() {

	e := 10
	f := "golang"
	g := 3.14
	h := true

	// Shorthand notated goes over VAR
	fmt.Printf("The type is: %T | The value is: %#v \n", a, a)
	fmt.Printf("The type is: %T | The value is: %#v \n", b, b)
	fmt.Printf("The type is: %T | The value is: %#v \n", c, c)
	fmt.Printf("The type is: %T | The value is: %#v \n", d, d)

	fmt.Println()

	fmt.Printf("The type is: %T | The value is: %v \n", e, e)
	fmt.Printf("The type is: %T | The value is: %v \n", f, f)
	fmt.Printf("The type is: %T | The value is: %v \n", g, g)
	fmt.Printf("The type is: %T | The value is: %v \n", h, h)
}
