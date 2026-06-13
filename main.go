//What is Go:
// Golang- an open source STATICALLY TYPED, COMPILED language. Developed to make
// software development at scale faster, simpler and reliable. 
// # It is simple- Only 25 keyowrds, No OOP inheritance. Easy to read, maintain and adopt
// # Built for Concurrency- Instead of using OS's threads and relying on OS to manage them, GO uses 
//Go Routines. GoRoutines are lightweight threads managed by the Go runtine and has size only of a few Kilobytes
// We can spin up thousands of routines concurrently, they communicate safetly using channels preventing data races
// # Fast compilation & execution, Robust and legendary standard library, static typing & garbage collector.

// COMMANDS:
// go mod init <moduleName> golang-fundamentals is the name of this module.
//go mod tidy - scans code and downloads missing externals pacakges, similar to npm i cmd in Node
//go run <fileName.go> or gun run . - compiles the code and executes it
//go build - complies the whole package and builds an very optimized binary executable
//go fmt - go's official code formatter (like prettier)
// go doc - to access whole golang documentation OFF-line.
//go run golang.org/x/tools/cmd/godoc@latest -http=:6060 - loads up documentation on browser port 6060

package main; // Tells Go this file is the main entry point for an application
// So go builds this file as an runnable executable

import "golang-fundamentals/golangconcepts"

func main(){
	golangconcepts.HelloWorld();
	golangconcepts.GolangDatatypes();
	golangconcepts.TypeConversion();
	golangconcepts.CompositeDataTypes();
	golangconcepts.SliceAddress();
	golangconcepts.GoStructs();
	golangconcepts.GoLoopsAndControlFlow();
	golangconcepts.GOFuncs(33, 44);
	golangconcepts.GoPointers();
	golangconcepts.MemoryAndGC();
	//gloangconcepts.GoInterfaces(); - read 09-interfaces.go
	golangconcepts.GoGenerics();
	//golangconepts.GoErrors(); - read 11-errors.go
	golangconcepts.GoConcurrency();
}