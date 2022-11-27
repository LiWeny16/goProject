package main

import (
	"fmt"
	"net/http"
	"os/exec"
	"time"
)

func Open(uri string) error {
	cmd := exec.Command("cmd", "/C", "start "+uri)
	return cmd.Run()
}

func main() {

	http.Handle("/", http.FileServer(http.Dir("./")))
	for i := 3; i >= 1; i-- {
		fmt.Printf("\n%d", i)
		time.Sleep(time.Duration(1) * time.Second)
	}
	fmt.Printf("\ngo-http服务器已经启动啦\n")
	fmt.Printf("连按Ctrl+c以关闭，直接×也可以关\n")
	fmt.Printf("v0.0.1\nOnion")
	time.Sleep(time.Duration(1) * time.Second)
	Open("http://127.0.0.1:3000")
	e := http.ListenAndServe(":3000", nil)

	fmt.Println(e)
}
