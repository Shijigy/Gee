package main

import (
	"fmt"
	"gee"
	"net/http"
)

func main() {
	// 使用 New() 创建gee的实例
	r := gee.New()
	// 使用 GET()方法添加路由（静态路由）
	r.GET("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	})
	r.GET("/hello", func(w http.ResponseWriter, req *http.Request) {
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})

	// 启动Web服务
	r.Run(":9999")
}
