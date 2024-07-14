package main

import (
	"fmt"

	"github.com/Lucioschenkel/fcutils-secret/pkg/events"
)

func main() {
	ed := events.NewEventDispatcher()
	fmt.Println(ed)
}
