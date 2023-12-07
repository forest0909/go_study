package main

import "net/http"

func handler(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("hello world"))
}
func main() {
	http.HandleFunc("/hello", handler)
	http.ListenAndServe(":8080", nil)
}
