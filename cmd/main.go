package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	// 获取程序运行时的目录
	var addr string
	var dir string
	flag.StringVar(&addr, "a", "0.0.0.0:80", "HTTP Server listen address")
	flag.StringVar(&dir, "d", "", "Hosting directory")
	flag.Parse()

	if dir == "" {
		path, err := os.Getwd()
		if err != nil {
			panic(err)
		} else {
			dir = path
		}
	}

	fmt.Printf("Server will be listen on %s", addr)
	// 静态文件处理器
	staticFileHandler := http.FileServer(http.Dir(dir))
	http.HandleFunc("/", staticFileHandler.ServeHTTP)
	// HTTP服务启动监听
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		panic(err)
	}
}
