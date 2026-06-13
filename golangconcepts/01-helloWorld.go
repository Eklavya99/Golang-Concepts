// 0-1. Hello world
package golangconcepts 

import "fmt" // Imports the Format package from standard library
//Fmt contains functions for formatting text and printing to console.

func HelloWorld() { 
	var msg string = "Hello world from Golang!!" // variable in go: (var name type = value)
	number := 111 // := can only be used inside functions. It automatically infers the type.
	const pi = 3.14159 // Cannot use := with consts, the value cannot be change bcz is locked in at compile time
	
	fmt.Println(msg, number, pi)
}