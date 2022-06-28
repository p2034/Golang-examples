package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// https://www.educative.io/answers/how-to-make-http-post-request-with-json-body-in-golang
// https://golangcode.com/get-http-method-from-request/
// https://pkg.go.dev/net/http#Request

type Data struct {
	Name string `json:"name"`
	Size uint64 `json:"size"`
}

func example_server() {
	http.HandleFunc("/example_post", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("/example_post")
		// check request
		if r.Method != http.MethodPost {
			/* set data in response (way 1) */
			io.WriteString(w, "This is not a POST request.")
		}

		/* get data from request (way 1) */
		// get from json
		var d Data
		body, _ := ioutil.ReadAll(r.Body)
		_ = json.Unmarshal(body, &d)
		//print
		fmt.Println(d.Name, d.Size)
	})

	http.HandleFunc("/example_get", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("/example_get")
		// check request
		if r.Method != http.MethodGet {
			/* set data in response (way 1) */
			io.WriteString(w, "This is not a GET request.")
		}

		/* set data in response (way 2) */
		// create json
		d := Data{"Image", 1024}
		body, _ := json.Marshal(d)
		// set data
		w.WriteHeader(http.StatusOK)
		w.Write(body)
	})

	// means localhost:8090
	http.ListenAndServe(":8090", nil)
}

func example_client() {
	// get

	// send request  and get answer
	resp, _ := http.Get("http://localhost:8090/example_get")
	defer resp.Body.Close()
	// read body
	body, _ := io.ReadAll(resp.Body)
	// unmarshal
	var d Data
	_ = json.Unmarshal(body, &d)
	// print
	fmt.Println(d.Name, d.Size)

	// post

	// send request  and get answer
	var d_for_sending = Data{"Film", 1024 * 1024}
	body, _ = json.Marshal(d_for_sending)
	_, _ = http.Post("http://localhost:8090/example_post", "application/json", bytes.NewBuffer(body))
	defer resp.Body.Close()
}

func main() {
	example_client()
}
