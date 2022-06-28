package main

import (
	"fmt"
	"time"
)

// https://www.educative.io/answers/what-is-a-goroutine
// https://www.educative.io/answers/what-are-channels-in-golang

func example_routine() {
	// start goroutine
	go func() {
		for {
			fmt.Print("1")
			time.Sleep(500 * time.Millisecond)
		}
	}()

	for {
		fmt.Println("3")
		time.Sleep(500 * time.Millisecond)
	}
}

func example_channel() {
	c := make(chan string)
	defer close(c)

	// start routine
	go func(c chan string) {
		for {
			name := <-c // receiving value from channel
			fmt.Println("Hello, ", name)
		}
	}(c)

	for {
		c <- "World"
		time.Sleep(500 * time.Millisecond)
	}
}

func example_select() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func(c chan<- string) {
		for {
			time.Sleep(time.Millisecond * 800)
			c <- "hey"
		}
	}(c1)

	go func(c chan<- string) {
		for {
			time.Sleep(time.Millisecond * 1600)
			c <- "hey"
		}
	}(c2)

	for {
		select {
		case res := <-c1:
			fmt.Println(res, "from 1")
		case res := <-c2:
			fmt.Println(res, "from 2")
		default:
			fmt.Println("nothing happen")
		}
		time.Sleep(time.Millisecond * 200)
	}
}

func main() {
	example_select()
}
