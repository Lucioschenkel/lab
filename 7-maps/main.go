package main

import "fmt"

func main() {
	salaries := map[string]int{
		"Wesley": 1000,
		"John":   2000,
		"Maria":  3000,
	}

	for name, salary := range salaries {
		fmt.Printf("O salario de %s eh %d\n", name, salary)
	}

	delete(salaries, "Wesley")

	fmt.Printf("%v\n", salaries)

	newSalaries := make(map[string]int)

	newSalaries["john"] = 100

	newSalaries2 := map[string]int{}

	fmt.Printf("%v\n", newSalaries2)
}
