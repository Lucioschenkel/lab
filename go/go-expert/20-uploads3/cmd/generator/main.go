package main

import (
	"fmt"
	"os"
)

func main() {
	i := 0
	for {
		f, err := os.Create(fmt.Sprintf("./tmp/file%d.txt", i))
		if err != nil {
			panic(err)
		}
		f.WriteString(fmt.Sprintf("Hello, file%d.txt", i))
		f.Close()
		i++
	}
}
