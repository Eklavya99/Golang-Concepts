package golangconcepts

// Memory management in go is very efficient and automated. Go handles allocation of and freeing up of memory
// How Go manages memory under hood:
// Go has two memory zones- Stack and heap
// The stack- Fast and temporay, Go uses it for routines, local variables and function args
// The heap- Shared & managed, Heap is large and central pool of memory for the whole program
// Memory on heap must be tracked and cleaned up by the garbage collector (GC) which takes time & CPU resources

// How does Go decide wheather to put variable in stack or heap?
// Go asks "Does this variable escape the scope of the function it was created in?"
// If No- it goes to stack i.e it is tempory variable
// if Yes- it goes to heap bcz some other part of our code might use it or need it.

func toStack() int {
	x := 42 // goes to stack because it is passed by value and once func finish executing variable x will be destroyed
	return x
}

func toHeap() *int {
	x := 42 // goes to heap because we are returning a pointer. Go looks at this variable x and sees 
	// "ok we are returning a pointer to address of some variable x,
	// it means others might use it so i should put it in heap bcz once the function finish executing this will be destroyed"
	return &x 
}

func MemoryAndGC(){
	
}