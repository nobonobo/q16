# Q16 Fixed-Point Arithmetic Library

## Overview
This project provides a **Q16.16 fixed-point arithmetic library** implemented in Go. It offers a lightweight alternative to floating‑point numbers for deterministic calculations, especially useful in embedded systems and graphics where performance and reproducibility matter.

## Features
- `Fixed` type representing a signed 32‑bit Q16.16 value.
- Basic arithmetic: addition, subtraction, multiplication, division.
- Mathematical helpers: trigonometric functions, exponentiation, rounding, and sign handling.
- Comprehensive unit tests covering all operations.

## Installation
```sh
# If this repository is a go module
go get ./...
```
Or, add the module path to your project’s `go.mod`.

## Usage Example
```go
package main

import (
    "fmt"
    "path/to/q16" // replace with the actual module import path
)

func main() {
    a := q16.FromFloat(1.5)   // convert float64 to Fixed
    b := q16.FromFloat(2.0)
    c := q16.Mul(a, b)        // multiply Fixed numbers
    fmt.Printf("Result: %f\n", q16.ToFloat(c))
}
```

## Running Tests
```sh
go test ./...
```
All tests should pass, confirming the correctness of the fixed‑point operations.

## License
MIT License – see the `LICENSE` file for details.

---
*This README is written in English as requested.*
