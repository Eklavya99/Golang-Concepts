//02-dataTypes.go

package golangconcepts;

import "fmt"

func GolangDatatypes() {
	var a int
	var b uint //unsigned
	// int8, int16, int32, int64, uint8, uint64
	var c byte
	var d float32 //float64
	var e bool    //T or F
	f := "" // composite type
	var complexNo complex64 = 10 + 11i // complex type
	var dollar rune = '$'
	

	// We also have composite type like- arrays, slices, strings & map

	fmt.Println(a, b, c, d, e, f, complexNo, dollar) // The variables have been defined but not declared, these are their default values
}

func TypeConversion(){
	var integerNumber int = 42;
	var decimalNumber float32 = 5.5;
	//WE cannot do: var total = integerNumber + decimalNumber XXX-Error-XXX
	
	var total = float32(integerNumber) + decimalNumber
	fmt.Println(total); 
}