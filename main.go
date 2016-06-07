package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
)

func echo(w http.ResponseWriter, r *http.Request) {
	text, _ := ioutil.ReadAll(r.Body)
	strText := string(text)
	fmt.Fprint(w, strText)
}
func echoBase64(w http.ResponseWriter, r *http.Request) {
	text, _ := ioutil.ReadAll(r.Body)
	strText := base64.StdEncoding.EncodeToString(text)
	fmt.Fprint(w, strText)
}
func main() {
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/echo_base64", echoBase64)
	http.ListenAndServe(":8192", nil)
}
