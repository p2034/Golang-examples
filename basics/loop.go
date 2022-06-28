package main

import (
	"fmt"
	"math/rand"
	"time"
)

func example_for_3_component() {
	sum := 0
	for i := 1; i < 5; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func example_for_while() {
	rand.Seed(time.Now().UnixNano())

	var i float32
	for i < 0.5 {
		i = rand.Float32()
		fmt.Println(i)
	}
}

func example_for_infinite() {
	for {
		fmt.Println("mmm...")
		break
	}
}

func example_for_each() {
	strings1 := []string{"hello", "world"}
	for i, s := range strings1 {
		fmt.Println(i, s)
	}

	strings2 := map[string]float32{"first": 1.1, "second": 2.2}
	for i, s := range strings2 {
		fmt.Println(i, s)
	}
}

func main() {
	example_for_each()
}
