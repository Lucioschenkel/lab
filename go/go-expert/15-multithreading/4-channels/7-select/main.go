package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type Message struct {
	ID  int64
	Msg string
}

func main() {
	c1 := make(chan Message)
	c2 := make(chan Message)
	var i int64 = 0

	go func() {
		for {
			time.Sleep(time.Second * 1)
			atomic.AddInt64(&i, 1)
			c1 <- Message{ID: i, Msg: "RabbitMQ"}
		}
	}()

	go func() {
		time.Sleep(time.Millisecond * 900)
		atomic.AddInt64(&i, 1)
		c2 <- Message{ID: i, Msg: "Kafka"}
	}()

	for {
		select {
		case msg1 := <-c1:
			fmt.Printf("received: %d %s\n", msg1.ID, msg1.Msg)
		case msg2 := <-c2:
			fmt.Printf("received: %d %s\n", msg2.ID, msg2.Msg)
		case <-time.After(time.Second * 4):
			println("timeout")
			// default:
			// 	println("default")
		}
	}

}
