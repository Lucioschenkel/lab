package main

import "fmt"

func main() {
	var myArray [3]int
	myArray[0] = 10
	myArray[1] = 20
	myArray[2] = 30

	for i, v := range myArray {
		fmt.Printf("O valor do indice eh %d e o valor eh %d\n", i, v)
	}
}
