package main

import (
	//"encoding/base64"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

const STR_LENGTH = 1000000
const N = 10
const M = 10

var CHARSET = make([]byte, 26)

func init() {
	rand.Seed(time.Now().Unix())
	for i := 0; i < 26; i++ {
		CHARSET[i] = byte('a' + i)
	}
}
func randStr(length int) string {
	data := make([]byte, length)
	for i := 0; i < length; i++ {
		data[i] = CHARSET[rand.Intn(26)]
	}
	return string(data)
}
func main() {
	chs := make([]chan int, N)
	for i := range chs {
		chs[i] = make(chan int)
	}
	start := time.Now()
	for i := 0; i < N; i++ {
		go func(ch chan int) {
			errCount := 0
			for i := 0; i < M; i++ {
				str := randStr(STR_LENGTH)
				//rsp, err := http.Post("http://127.0.0.1:8192/echo_base64", "text/plain", strings.NewReader(str))
				rsp, err := http.Post("http://127.0.0.1:8192/echo", "text/plain", strings.NewReader(str))
				if err != nil {
					fmt.Println(err)
					errCount++
					continue
				}
				rspStr, err := ioutil.ReadAll(rsp.Body)
				if err != nil {
					fmt.Println(err)
					errCount++
					continue
				}
				//ansStr := base64.StdEncoding.EncodeToString([]byte(str))
				ansStr := str
				if string(rspStr) != ansStr {
					errCount++
					continue
				}
			}
			ch <- errCount
			close(ch)
		}(chs[i])
	}
	errCount := 0
	for i := 0; i < N; i++ {
		errCount += <-chs[i]
	}
	stop := time.Now()
	size := N * M * STR_LENGTH / 1024.0 / 1024.0
	exeTime := stop.Sub(start).Seconds()
	fmt.Printf("Total length: %.2f MB\n", size)
	fmt.Printf("Time: %.2f s\n", exeTime)
	fmt.Printf("Speed: %.2f MB/s\n", size/exeTime)
	fmt.Printf("Error count: %d\n", errCount)
}
