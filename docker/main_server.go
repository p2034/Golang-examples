package main

import (
	"fmt"
	"net/http"
)

func mainloop() {
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Hello from server!</h1>")
	})

	// means localhost:8090
	http.ListenAndServe(":8090", nil)
}

func main() {
	mainloop()
}
