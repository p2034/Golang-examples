package main

import (
	"fmt"
	"reflect"
)

func example_casting() {
	var a float32 = 1.1

	fmt.Println(reflect.TypeOf(int(a))) // int
	fmt.Println(int(a))                 // 1
}

func example_assertion() {
	var value interface{} = "test"

	var value1 string = value.(string)
	fmt.Println(value1) // test

	/*
		var value2 int = value.(int) // panic
		fmt.Println(value2)
	*/

	value2, check := value.(int)
	if check {
		fmt.Println(value2)
	} else {
		fmt.Println("Not int")
	}
}

func example_switches_helper(i interface{}) {
	switch v := i.(type) { // 'type' is a keyword
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func example_switches() {
	example_switches_helper(21)
	example_switches_helper("hello")
	example_switches_helper(true)
}

func main() {
	example_switches()
}
