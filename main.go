package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/calc", func(w http.ResponseWriter, r *http.Request) {
		aStr := r.URL.Query().Get("a")
		a, _ := strconv.Atoi(aStr)
		bStr := r.URL.Query().Get("b")
		b, _ := strconv.Atoi(bStr)
		fmt.Fprintf(w, "%v\n", a+b)
	})

	if err := http.ListenAndServe(":80", nil); err != nil {
		panic(err)
	}
}
