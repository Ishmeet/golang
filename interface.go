package main

import (
	"fmt"
)

func init() {
	fmt.Println("Init func invoked")
}

// Answer 2
type myMethod interface {
	swap() (int, int)
}

// Answer 1
func (s swap1) swap() (int, int) {
	return s.b, s.a
}

// Answer 1
func (s swap2) swap() (int, int) {
	return s.b, s.a
}

type swap1 struct {
	a, b int
}

type swap2 struct {
	a, b int
}

func main() {
	s1 := &swap1{a: 1, b: 2}
	a, b := s1.swap()
	fmt.Println(a, b)
	s2 := &swap2{a: 3, b: 4}
	a, b = s2.swap()
	fmt.Println(a, b)
}
