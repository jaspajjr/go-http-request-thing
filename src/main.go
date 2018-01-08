package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func adder(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func fooPageHandler(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.Query()
	foo := queryString["foo"]
	bar, _ := strconv.Atoi(foo[0])
	total := adder(bar)
	fmt.Fprintln(w, "This is the foo page")
	fmt.Fprintln(w, total)
}

func indexPageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is the home page")
}

func aboutPageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is the about page")
}

func loggerMiddleware(nextHandler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Request recieved: %v\n", r)
		nextHandler.ServeHTTP(w, r)
		fmt.Println("Request handled successfully")
	})
}

func main() {
	fmt.Println("It's running and what not.")

	mux := http.NewServeMux()

	mux.Handle("/", loggerMiddleware(http.HandlerFunc(indexPageHandler)))
	mux.Handle("/about", loggerMiddleware(http.HandlerFunc(aboutPageHandler)))
	mux.Handle("/foo", loggerMiddleware(http.HandlerFunc(fooPageHandler)))

	http.ListenAndServe(":8080", mux)
}
