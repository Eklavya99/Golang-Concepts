package golangconcepts

import "fmt"

// Go only has for loop.
func GoLoopsAndControlFlow(){
	//simple loop
	for i := 0; i < 5; i++{
		fmt.Println(i)
	}
	
	// Looping over collections
	skus := []string{"Go-book", "Chip-set", "Gopher-toy"}
	for index, value := range skus{
		fmt.Printf("Index: %d, SKI: %s\n", index, value)
	}

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 7, 8, 0}
	for _, num := range numbers{
		if num == 0 {
			fmt.Printf("Number is 0")
			break;
		} else if num%2 == 0{
			fmt.Printf("Number %d is even", num)
		}else{
			continue
		}
	}

	var str string = "Abcdefghijklmopqrst"

	for i := 0; i < len(str); i++{
		fmt.Println(str[i]) // return bytes
	}
	for i, s := range str{
		fmt.Println(i, s); // returns index and character/rune
	}

	kv := map[string]int{
		"Abc": 1,
		"def": 2,
		"hij": 3,
	}

	for k, v := range kv{
		fmt.Println(k, v);
	}
}