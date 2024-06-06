package main

import (
	"fmt"
	"github.com/Lucioschenkel/goexpert/12-packaging/1/math"
)

func main() {
  m := math.Math{A: 1, B: 2}
  fmt.Println(m.Add())
}
