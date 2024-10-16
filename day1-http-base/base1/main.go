package main

import (
	"fmt"
	"log"
	"net/http"
)

//$ curl http://localhost:9999/
//URL.Path = "/"
//
//$ curl http://localhost:9999/hello
//Header["Accept"] = ["*/*"]
//Header["User-Agent"] = ["curl/7.54.0"]

func main() {
	// 设置2个路由 `/` 和 `/hello` 分别绑定 indexHandler 和 helloHandler
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	// 启动Web服务
	log.Fatal(http.ListenAndServe(":9999", nil))
}

// handler echoes r.URL.Path
func indexHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
}

// handler echoes r.URL.Header
func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}
