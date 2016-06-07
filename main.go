package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func echo(w http.ResponseWriter, r *http.Request) {
	text, _ := ioutil.ReadAll(r.Body)
	strText := string(text)
	fmt.Fprint(w, strText)
}
func main() {
	http.HandleFunc("/echo", echo)
	http.ListenAndServe(":8192", nil)
}
