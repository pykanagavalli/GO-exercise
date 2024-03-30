package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hi")
	})
	port := 8500
	fmt.Println("server running on ", +port)
	http.ListenAndServe(":8500",nil)

}
		