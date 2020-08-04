package main

import "fmt"

func init() {

}

func swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

func routine(n int, msg chan string) {
	for i := 0; i < 5; i++ {
		fmt.Println("Hello", i+1)
	}
	msg <- "some string"
}

type platform interface {
	send()
}

type config1 struct {
	endpoint string
}

type config2 struct {
	opsendpoint string
}

func (c *config1) send() {
	fmt.Println(c.endpoint)
}

func (c *config2) send() {
	fmt.Println(c.opsendpoint)
}

func main() {
	a, b := 1, 2
	msg := make(chan string)
	swap(&a, &b)
	fmt.Println(a, b)
	go routine(5, msg)
	fmt.Println(<-msg)

	c1 := &config1{endpoint: "https://directlink.com"}
	c2 := &config2{opsendpoint: "https://internal.directlink.com"}
	c1.send()
	c2.send()
}
