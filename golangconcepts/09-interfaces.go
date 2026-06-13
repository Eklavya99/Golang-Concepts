package golangconcepts

import "fmt"

// In C#, a class must explicitly state that it implements an interface: class MyCSharpClass : IMyCSharpInterface
// In Go, interfaces are satisfied implicitly. If a type has methods defined in an interface, then it implements it
// In Go, an interface is a type defined by a set of method signatures.


type Speaker interface {
	Speak() string
}

type Dog struct{}

func (d Dog) Speak() string { 
// type Dog has implicity implemented the interface- if a type has methods defined in an interface,
//  it implements it. SIMPLE.
	return "Woof Woof!"
}

// In other languages like C#, you need to state it explicitly that a type(or class) implements an interface
// ex- public class Dog : ISpeaker {...}

// An empty interface specifies zero methods. Because it has no requirements, every single type in Go satisfies it.
func PrintAnything(value interface{}){
	fmt.Println(value)
}
// In modern Go, you'll see the alias any used instead of interface{}, but they are identical under the hood.
func PrintAny(value any){
	fmt.Println(value)
}

// Go doesn't have class inheritance, so it uses composition. 
// You can create complex interfaces by embedding simpler ones inside them.
type Reader interface {
	Read() string
}
type Writer interface{
	Write() string
}
type ReaderAndWriter interface{
	Reader
	Writer
	// Any type that implements Read() and Write() methods automatically satisfies ReaderAndWriter interface
}
