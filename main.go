package main

import (
	"fmt"
	"net/http"
	"os"
)

func param(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "hello world")
	header := req.Header

	// fmt.Fprintln(res, "Header全部数据:", header)

	var acc []string = header["Accept"]
	for _, n := range acc {
		fmt.Fprintln(res, "Accepth内容:", n)
	}

	environ := os.Environ()
	for i := range environ {
		fmt.Fprintln(res, environ[i])
	}

}

func main() {
	server := http.Server{
		Addr: "localhost:8090",
	}
	http.HandleFunc("/param", param)
	server.ListenAndServe()
}
