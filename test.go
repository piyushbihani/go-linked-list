package main

import (
	"fmt"
)

type list1 struct{
	data int
	next *list1
}

func add_node(data int)*list1{
	node := new(list1)
	node.data = data
	node.next = nil
	return node
}

func (node list1) print_node(){
	fmt.Print("\nNode: ")
	for i := 0; i < 2; i++ {
		fmt.Printf("%d -> ", node.data)
		if node.next != nil{
			node = *node.next
		}else{
			break
		}
	}
	fmt.Print("NULL\n")
}

func test(){
	head := add_node(10)
	head2 := add_node(20)
	head.next = head2
	head.print_node()

}