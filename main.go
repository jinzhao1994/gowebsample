package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

var port int

func init() {
	flag.IntVar(&port, "port", 8192, "")
	flag.Parse()
}

func writeMeta(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Got Request:", r)
	fmt.Fprintln(w, r.Method, r.URL)
	fmt.Fprintln(w)
	fmt.Fprintln(w, "Header:")
	for k, v := range r.Header {
		fmt.Fprintln(w, k, ":", v)
	}
	fmt.Fprintln(w)
}

func echo(w http.ResponseWriter, r *http.Request) {
	writeMeta(w, r)
	text, _ := ioutil.ReadAll(r.Body)
	strText := string(text)
	fmt.Fprintln(w, "Body:")
	fmt.Fprint(w, strText)
}
func echoBase64(w http.ResponseWriter, r *http.Request) {
	writeMeta(w, r)
	text, _ := ioutil.ReadAll(r.Body)
	strText := base64.StdEncoding.EncodeToString(text)
	fmt.Fprintln(w, "Body:")
	fmt.Fprint(w, strText)
}

func main() {
	fmt.Println("Listen on port", port)
	http.HandleFunc("/", echo)
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/echo_base64", echoBase64)
	http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
