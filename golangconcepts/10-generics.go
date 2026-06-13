package golangconcepts

import "fmt"

// Introduced in Go 1.18, generics bring type parameters to the language.
// Before generics, if you wanted to write a function that returned the minimum of two numbers,
// you had to write separate functions for every single type- int, float, uint,
// or use the empty interface (any). And Using 'any' meant risking compile-time type safety
// Generics solve this by allowing you to pass "types as arguments" when defining functions or structures.

//T is a type parameter and can be anything - int, uint, string...
//Comparable is a type constraint it means types which support == , !=, >= etc type operation 
// i.e comparison operation
func findIndex[T comparable](slice []T, target T) int{
	for i, item := range slice{
		if item == target{
			return i
		}
	}
	return -1
}

// Generic types & interfaces
type Box[T any] struct {
	Content T
}

// Type contraints and inference
// Type constraints restrict what types can be passed as T
type Numeric interface { // we defined our own numeric contraint.
	~int | ~float64 
}
func addNums[T Numeric](a T, b T) T{
	return a + b
}

func GoGenerics(){
	ints := []int{1, 3, 5, 66, 4, 2, 88, 9}
	fmt.Println(findIndex[int](ints, 66))

	strs := []string{"Bob", "Noddy", "Tom", "Jerry"}
	fmt.Println(findIndex(strs, "Oggy")) // Go automatically infers the type in case if not explicitly mentioned
}