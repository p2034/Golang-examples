package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

// just working with map
func example_map() {
	// creating map
	s := map[string]interface{}{"data": []string{"a", "b"}, "file": "some text"}

	// marshal or json
	json_test, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Marshaled value:", string(json_test))

	// unmarshal our map
	err = json.Unmarshal(json_test, &s)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Unmarshaled value:", s["file"])
}

// marshal structures
func example_struct() {
	type pet struct {
		Name   string `json:"name"`
		Age    int    `json:"age"`
		Animal string `json:"animal"`
	}

	// creating map
	s := pet{"star", 6, "dog"}

	// marshal or json
	json_test, err := json.Marshal(s)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("Marshaled value:", string(json_test))
}

// reading and writing file struct person
func example_file() {
	type person struct {
		Name   string  `json:"name"`
		Age    int     `json:"age"`
		Height float32 `json:"height"`
	}

	data := person{}

	file, _ := ioutil.ReadFile("example/json.json")
	_ = json.Unmarshal([]byte(file), &data)

	fmt.Println(data)
	data.Age = 25

	file, _ = json.Marshal(data)
	_ = ioutil.WriteFile("example/json.json", file, 0644)
}

func example_decoder() {
	type Message struct {
		Name, Text string
	}
	sstr := `{"Name":"Alex","Text":"mmm"}
	{"Name":"Jhon","Text":"what?"}
	{"Name":"Violet","Text":"nothing"}
	{"Name":"Felix","Text":"sure?"}`

	dec := json.NewDecoder(strings.NewReader(sstr))
	for {
		var m Message
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: %s\n", m.Name, m.Text)
	}
}

func main() {
	example_decoder()
}
