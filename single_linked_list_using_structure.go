
package main

import (
	"fmt"
)

type singly_node struct{
	data int
	next *singly_node
}

var singly_size uint16 = 0

func (head *singly_node)new_node_at_start(data int)*singly_node{
	node_to_add := new_node(data)
	node_to_add.next = head
	head = node_to_add
	return head
}

func (head *singly_node)new_node_at_middle(data int, position int)*singly_node{
	if position > int(singly_size) + 1 || position <= 0{
		fmt.Print("Invalid Position\n")
		return head
	}else if position == 1{
		return head.new_node_at_start(data)
	}
	node_to_add := new_node(data)
	curr := head
	for i := 1; i < position - 1; i++{
		curr = curr.next
	}
	temp := curr.next
	curr.next = node_to_add
	node_to_add.next = temp
	return head
}

func new_node(data int)*singly_node{
	singly_size += 1
	node := new(singly_node)
	node.data = data
	node.next = nil
	return node
}

func (head *singly_node)delete_at_first()*singly_node{
	if singly_size <= 0{
		fmt.Print("List is empty\n")
		return head
	}
	singly_size--
	if singly_size == 0 || head.next == nil{
		return nil
	}
	return head.next
}

func (head *singly_node)delete_at_middle(position int)*singly_node{
	if singly_size <= 0 || position <= 0 || position > int(singly_size){
		fmt.Print("Invalid\n")
		return head
	}
	if position == 1{
		return head.delete_at_first()
	}
	temp := head
	for i := 1; i < position - 1 && temp.next != nil; i++{
		temp = temp.next
	}
	singly_size--
	if temp.next.next != nil{
		temp.next = temp.next.next
	}else{
		temp.next = nil
	}
	return head
}

func (head *singly_node)delete_at_last()*singly_node{
	if singly_size <= 0{
		fmt.Print("List is empty\n")
		return head
	}
	temp := head
	for ; temp.next != nil && temp.next.next != nil; temp=temp.next{
	}
	temp.next = nil
	singly_size--
	if singly_size == 0{
		return nil
	}
	return head
}

func (head *singly_node)print_list(){
	fmt.Print("\nNode: ")
	for temp := head; temp != nil; temp = temp.next{
		fmt.Printf("%d -> ", temp.data)
	}
	fmt.Print("NULL\n")
}


func single_linked_list(){
	var head *singly_node = nil
	temp := head
	for{
		fmt.Print("\nEnter your choice: ")
		var input uint
		fmt.Scanf("%d\n", &input)
		if input == 1{
			var data int
			fmt.Print("Enter the value of node: ")
			fmt.Scanf("%d\n", &data)
			if head == nil{
				head = new_node(data)
				temp = head
			}else{
				node := new_node(data)
				temp.next = node
				temp = node
			}
		}else if input == 2{
			var data int
			fmt.Print("Enter the value of node: ")
			fmt.Scanf("%d\n", &data)
			head = head.new_node_at_start(data)
		}else if input == 3{
			var data, pos int
			fmt.Print("Enter the value of node and position: ")
			fmt.Scanf("%d %d\n", &data, &pos)
			head = head.new_node_at_middle(data, pos)
		}else if input == 4{
			head = head.delete_at_last()
		}else if input == 5{
			head = head.delete_at_first()
		}else if input == 6{
			var pos int
			fmt.Print("Enter the position: ")
			fmt.Scanf("%d", &pos)
			head = head.delete_at_middle(pos)
		}else if input == 7{
			fmt.Printf("singly_singly_sin = %d\n", singly_size)
		}else if input == 8{
			head.print_list()
		}
	}
}