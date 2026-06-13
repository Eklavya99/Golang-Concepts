# Golang Fundamentals

A hands-on Go learning project with small, runnable examples that cover core language concepts — from variables and types to concurrency.

## Prerequisites

- [Go](https://go.dev/dl/) 1.24.2 or later

## Getting Started

Clone the repo, then run everything from the project root:

```bash
go run .
```

This runs `main.go`, which calls each concept in order.

To run a single example file directly:

```bash
go run golangconcepts/01-helloWorld.go
```

> Note: Individual files in `golangconcepts/` are meant to be studied as part of the package. Running them alone may not work unless they include a `main` function.

## Project Structure

```
golang-fundamentals/
├── main.go                 # Entry point — runs all examples
├── go.mod
└── golangconcepts/         # One file per topic
    ├── 01-helloWorld.go
    ├── 02-dataTypes.go
    ├── 03-compositeDatatypes.go
    ├── 04-structs.go
    ├── 05-loopsAndControlFlow.go
    ├── 06-functions.go
    ├── 07-pointers.go
    ├── 08-memoryAndGC.go
    ├── 09-interfaces.go
    ├── 10-generics.go
    ├── 11-errors.go
    └── 12-GoConcurrency.go
```

## Topics Covered

| # | File | What you'll learn |
|---|------|-------------------|
| 01 | `01-helloWorld.go` | Variables, constants, `:=` shorthand, `fmt.Println` |
| 02 | `02-dataTypes.go` | Primitive types, type conversion |
| 03 | `03-compositeDatatypes.go` | Arrays, slices, maps, strings |
| 04 | `04-structs.go` | Structs, embedding, methods |
| 05 | `05-loopsAndControlFlow.go` | `for` loops, `range`, `if`/`else`, `break`/`continue` |
| 06 | `06-functions.go` | Functions, multiple return values, variadic functions |
| 07 | `07-pointers.go` | Pointers, value vs pointer receivers |
| 08 | `08-memoryAndGC.go` | Stack vs heap, escape analysis, garbage collection |
| 09 | `09-interfaces.go` | Implicit interfaces, `any` / `interface{}` |
| 10 | `10-generics.go` | Type parameters, constraints, generic structs |
| 11 | `11-errors.go` | Error values, the `error` interface, `panic`/`recover` |
| 12 | `12-GoConcurrency.go` | Goroutines, channels (buffered & unbuffered) |

## How to Study

1. Start with `main.go` for a high-level overview of Go and useful commands.
2. Open each file in `golangconcepts/` in order — every file has inline comments explaining the concepts.
3. Run `go run .` after reading a section to see the output.
4. Two topics are commented out in `main.go` and meant to be read on their own:
   - **Interfaces** — see `09-interfaces.go`
   - **Errors** — see `11-errors.go`

## Useful Go Commands

```bash
go mod init <module-name>   # Initialize a new module
go mod tidy                 # Download and clean up dependencies
go run .                    # Compile and run the project
go build                    # Build an executable binary
go fmt ./...                # Format all Go files
go doc                      # Browse standard library docs offline
```

To browse Go docs in a browser:

```bash
go run golang.org/x/tools/cmd/godoc@latest -http=:6060
```

Then open [http://localhost:6060](http://localhost:6060).

## Why Go?

Go is a statically typed, compiled language built for simplicity and concurrency. Key strengths:

- **Simple syntax** — 25 keywords, no class inheritance
- **Built-in concurrency** — lightweight goroutines and channels
- **Fast compilation** — produces optimized binaries
- **Strong standard library** — batteries included for common tasks
