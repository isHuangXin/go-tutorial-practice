package main

import (
	"fmt"
	"math"
)

// what is export name?
// only the first character is capital, the module can be cited.
func main() {
	fmt.Println("export name math.Pi: ", math.Pi)
	fmt.Println("export name math.pi: ", math.pi)
}
