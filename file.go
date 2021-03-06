package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

// https://zetcode.com/golang/file/

/*
	READING
*/

/*
Creating files and they removing worked with os module
os.Stat("words.txt")
file, err := os.Create("empty.txt")
err := os.Remove("words.txt")
*/
func example_stat() {
	fInfo, err := os.Stat("example_data/text_example.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(fInfo)
}

// reading full file with ioutil
func example_full_reading() {
	file, _ := ioutil.ReadFile("example/text.txt")

	fmt.Println(string(file))
}

// reading line by line with os and bufio
func example_line_by_line_reading() {
	f, _ := os.Open("example/text.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
		time.Sleep(100 * time.Millisecond)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

/*
	WRITING
*/

// writing full file with ioutil
func example_full_writing() {
	val := "Hello world1!\nHello world2!\nHello world3!\nHello world4!\n"
	data := []byte(val)

	err := ioutil.WriteFile("example/text.txt", data, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

// append file with os
func example_append_writing() {
	f, err := os.OpenFile("example/text.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if _, err := f.WriteString("Hello world5!\n"); err != nil {
		log.Fatal(err)
	}
}

func renderFile(w http.ResponseWriter, filename string) {
	fmt.Println("Read request: " + filename)
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Cann't open file: " + filename)
	} else {
		w.Write(file)
	}
}

//	like it was from user request
func example_image() {
	http.HandleFunc("/image", func(w http.ResponseWriter, r *http.Request) {
		renderFile(w, "./home.jpeg")
	})

	http.ListenAndServe(":8090", nil)
}

func main() {
	example_image()
}
