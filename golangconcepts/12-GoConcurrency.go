package golangconcepts

import (
	"fmt"
	"time"
)

// Go replaces complex OS thread management with a beautiful abstraction layer.
// A goroutine is a lightweight thread of execution managed by the Go runtime.
// OS thread typically has size of 1-2 MB while go-routines have size of 2 KBs only
// To start one, you simply prepend the keyword 'go' to any function call.
// The main function runs in its own goroutine.
// If main finishes, the program exits immediately, killing all other background goroutines 
// even if they aren't done.

func printNumbers(){
	for i := 1; i <= 3; i++{
		time.Sleep(100 * time.Millisecond)
		fmt.Println(i)
	}
}

// Channels are pipelines that connect concurrent routines, we can send data from one routine
// to another one, handling synchronization naturally without explicit locks
// un-buffered channels- (synchronous) sending data blocks the sender until a recevier gets it
// receiving the data then blocks the recevier until the data is sent.
// buffered channels- (asynchronous) buffered channels have a pre-defined capacity, sending data
// only blocks the sender when the buffer is full, receiving only block when the buffer is 
// empty

func GoConcurrency(){
	go printNumbers() // starts concurrently
	fmt.Println("main(..) is running")
	time.Sleep(500 * time.Millisecond) // wait so main() doesnt exit instantly

	ch := make(chan string) // unbuf
	//bufCh := make(chan int, 3) // buf

	ch <- "Hello" // send data to channel
	msg := <-ch // get data from channel
	fmt.Println(msg)
}