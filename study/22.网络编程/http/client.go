package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	client := new(http.Client)
	req, _ := http.NewRequest("GET", "http://localhost:8080/hello", nil)
	res, _ := client.Do(req)
	b, _ := io.ReadAll(res.Body)
	fmt.Println(string(b))
}
