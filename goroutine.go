package main

import (
	"fmt"
)

func init() {

}

func main() {
	msg := make(chan string)

	go func(n int) {
		for i := 0; i < n; i++ {
			fmt.Println("hello world")
		}
		msg <- "done"
	}(5)
	<-msg
}
