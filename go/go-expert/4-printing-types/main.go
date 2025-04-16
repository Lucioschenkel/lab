package main

import (
	"fmt"
)

type ID int

var (
	e float64 = 1.2
	f ID      = 1
)

func main() {
	fmt.Printf("O tipo de E eh: %T", e)
	fmt.Printf("O tipo de F eh: %T", f)
}
