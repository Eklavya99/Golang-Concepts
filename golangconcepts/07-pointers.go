package golangconcepts

import "fmt"

// at simplest, Pointers are just variables which store memory address of another variable
// Go provides two primary operators to work with pointers
// & (address of operator)
// * (de-reference operator)

type Player struct{
	Name string
	Score int
}
func levelUp(p *Player){
	p.Score += 100
}

func (p Player) powerUp(){ // Operates on a copy of Player struct, Also called value receivers
		p.Score += 10;
	}
func (p *Player) skillUp(){ // operates on the orignal Players struct, Also called pointer receivers
	p.Score += 5
} 

func GoPointers(){
	age := 25
	var ptr *int = &age
	//&age gives the address of variable age
	//*int tells go this variable 'ptr' holds address of an int type variable
	fmt.Println("Address of age: ", &age) // Outputs 0xc00f435e
	fmt.Println("value store in ptr: ", ptr) // Outputs 0xc00f435e, same as above

	p1 := Player{Name: "Drox", Score: 350,}
	levelUp(&p1) //passing address of p1
	fmt.Println(p1.Score)
}