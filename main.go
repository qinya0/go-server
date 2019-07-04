package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	testhttp()
}

func testhttp() {
	fmt.Println("start HandleFunc")
	http.HandleFunc("/", ServeHTTP)
	fmt.Println("start ListenAndServe")
	http.ListenAndServe("0.0.0.0:8080", nil)
	fmt.Println("end ListenAndServe")
}

func ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// 按json转
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read all err: %s", err.Error())
		return
	}

	fmt.Printf("method:%s", r.Method)
	fmt.Printf("body:%s", string(b))

	rw.WriteHeader(http.StatusOK)
}
