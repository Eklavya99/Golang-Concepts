package golangconcepts

import (
	"errors"
	"fmt"
)

// Go does not have exceptions i.e no try-catch-finally blocks for standard control flow.
// Instead, errors are just values returned from functions like any other variable.

func divideNums(a int, b int) (int, error) {
	if (b == 0){
		return 0, errors.New("Cannot divide a number by 0") 
	}
	return a/b, nil
}

func GoErrors(){
	result, err := divideNums(10, 0)
	if err != nil {
		fmt.Println("An error has occured: ", err)
		return
	}
	fmt.Println("result: ", result)
}

// Go's built-in error interface:
type error interface{
	Error() string // anything that has this method automatically implements the error interface
}

type DivideError struct{
	n1 int
	n2 int
	Err error
}
func (e DivideError) Error() string{
	return fmt.Sprintf("Division invalid: %v", e.Err)
}

// Two ways to spin up an error-
// errors.New- for basic static text i.e when no dynamic data is needed for err, Like user not found 
// fmt.Errorf- used when you need to format dyanamic data into the err
var err1 error = errors.New("An error occurred")
var err2 error = fmt.Errorf("failed to connect to host %s: %w", "localhost", err1) ///wrapping errors

