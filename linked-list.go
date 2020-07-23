package main

import (
	"fmt"
)

func init() {

}

type node struct {
	data int
	next *node
}

func print(head *node) {
	curr := head
	for {
		if curr != nil {
			fmt.Println(curr.data)
			curr = curr.next
		} else {
			break
		}
	}
}

func reverse(head **node) {
	curr := *head
	prev := *head
	prev = nil
	next := *head
	next = nil
	for {
		if curr != nil {
			next = curr.next
			curr.next = prev
			prev = curr
			curr = next
		} else {
			break
		}
	}
	*head = prev
}

func main() {
	head := &node{data: 1,
		next: &node{
			data: 2,
			next: &node{
				data: 3,
				next: &node{
					data: 4,
					next: &node{
						data: 5,
						next: nil,
					},
				},
			},
		}}
	print(head)
	reverse(&head)
	print(head)

}
