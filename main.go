package main

import (
	"fmt"
	"time"
)

var flag bool

func writer() {
	for ;; {
		flag = !flag
		time.Sleep(1000 * time.Millisecond)
	}
}

func reader() {
	for ;; {
		fmt.Println("flag: ", flag)
		time.Sleep(1000 * time.Millisecond)
	}
}

func main() {
	flag = true
	go writer()
	reader()
}
