package main

import "github.com/Lucioschenkel/goexpert/12-packaging/3/math"
// go mod edit -replace github github.com/Lucioschenkel/goexpert/12-packaging/3/math=../math

func main() {
  m := math.Math{
    A: 1,
    B: 2,
  }
  println(m.Add())
}
