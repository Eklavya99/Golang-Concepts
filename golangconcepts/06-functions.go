package golangconcepts

import (
	"errors"
	"fmt"
)

func add(x int, y int, z int) int { // func funcName(params) return-type{...logic...}
	return x + y + z
}

func divide(numerator float64, denominator float64) (float64, error) {
	if denominator == 0 {
		return 0, errors.New("Cannot divide by 0")
	}
	return numerator/denominator, nil
}

// Varidac functions
func multiply(numbers ...int) int { // varidac functions can take any number of args
	product := 1
	for _, ints := range numbers{
		product *= ints
	}
	return product
}

func GOFuncs(x int, y int) (int, float64, error) { // go functions can have multiple return types
	res1 := add(x, y, 0)
	res2, err := divide(float64(x), float64(y))
	if err != nil{
		return 0, 0, err
	}
	res3 := multiply(2, 4, 6, 8)
	res4, res5 := NamedReturn()
	fmt.Println(res3, res4, res5);
	return res1, res2, nil
}

// Named return types
func NamedReturn() (width int, height int){
	width = 1920
	height = 1080
	return
}