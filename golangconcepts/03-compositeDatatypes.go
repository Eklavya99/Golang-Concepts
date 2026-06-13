package golangconcepts

import "fmt"

func CompositeDataTypes() {
	// Primitive (int, string, bool etc) - Holds a single value
	// Composite have multiple values

	// Arrays- FIXED SIZE as define during initialization
	var numbers [3]int // an array numbers of size 3 holding ONLY int data type
	names := [5]string{"Alice", "Bob", "Mickey", "Harry", "Ramandeep"}
	fmt.Println(names[2], numbers[1])

	// Slices- DYNAMIC SIZE
	var primes []int
	primes = append(primes, 2, 3, 5)

	superheros := make([]string, 3, 10) // make(type, len, cap)
	superheros = append(superheros, "batman", "superman", "spiderman", "hulk", "Ironman", "flash")
	println(superheros[2:5])
	println(superheros[:2])

	// Maps- Key-value pairs
	userRoles := make(map[string]string)
	userRoles["User-Eklavya"] = "admin"
	userRoles["User-unknown"] = "viewer"

	ports := map[string]int{
		"httpPort": 80,
		"httpsPort": 443,
	}
	println(ports["httpPort"])

	// Struct - Composite type to hold multiple related data together
	type Order struct{
		Id int
		OrderItems []string
		TotalCost float64
	}
	type User struct{
		Id int
		Username string
		Order []Order
		UserSettings map[string]string
	}
	u1Orders := []Order{
		{
			Id: 3924,
			TotalCost: 113.32,
		}, {
			Id: 9020,
			TotalCost: 554.12,
		},
	}
	u1 := User{
		Id: 101,
		Username: "HokeyPokey111",
		Order: u1Orders,		
	}
	println(u1.Username)
}

func SliceAddress() {
    // Length 0, Capacity 2
    s := make([]int, 0, 2) 
    
    s = append(s, 10)
    s = append(s, 20)
    fmt.Printf("Before: len=%d cap=%d ptr=%p\n", len(s), cap(s), s)
    
    // This append forces a reallocation because capacity is full!
    s = append(s, 30)
    fmt.Printf("After:  len=%d cap=%d ptr=%p\n", len(s), cap(s), s)
}