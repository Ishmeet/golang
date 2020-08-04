package main

import "fmt"

func duplicateOrMissing(list []int) ([]int, []int) {
	m := make(map[int]int)
	for _, v := range list {
		m[v]++
	}
	dupList := []int{}
	misList := []int{}
	for i, l := range m {
		if l > 1 {
			dupList = append(dupList, i)
		} else if l < 1 {
			misList = append(misList, i)
		}
	}
	return dupList, misList
}

func main() {
	a := []int{1, 2, 3, 4, 4, 5, 6, 6, 8}
	fmt.Println(duplicateOrMissing(a))
	f := 10.11
	fmt.Println(3 * int(f))
	fmt.Println([]byte("ABCD"))
}
