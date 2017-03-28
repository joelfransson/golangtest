package main

import (
	"fmt"
	"time"
)

func high(c chan int) {
	for i := 500; i < 600; i++ {
		c <- i
	}
}

func low(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
}

func channelFun() {
	fmt.Println("start")
	c1 := make(chan int)
	c2 := make(chan int)

	go low(c1)
	go high(c2)

	for i := 0; ; i++ {
		select {
		case n := <-c1:
			fmt.Println(fmt.Sprintf("%v", n))
			break
		case n := <-c2:
			fmt.Println(fmt.Sprintf("%v", n))
			break
		}
		time.Sleep(10 * time.Millisecond)
		if i > 20 {
			break
		}
	}
}
