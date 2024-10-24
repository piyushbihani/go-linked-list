
package main

import (
	"fmt"
	"os"
)

type doubly_node struct{
	data int
	next, prev *doubly_node
}

var doubly_size uint = 0

func (front *doubly_node)add_at_start(data int)*doubly_node{
	node := doubly_new_node(data)
	if doubly_size == 0{
		return node 
	}
	node.next = front
	if front != nil{
		front.prev = node
	}
	front = node
	return front
}

func (front *doubly_node)add_at_middle(data int, position int, back *doubly_node)(*doubly_node, *doubly_node){
	if position > int(doubly_size) + 1 || position <= 0{
		fmt.Print("Invalid Position\n")
		return front, back
	}else if doubly_size == 0 || position == 1{
		front = front.add_at_start(data)
		if doubly_size == 1{
			back = front
		}
		return front, back
	}
	curr := front
	prev_node := curr
	for i := 1; i < position; i++{
		prev_node = curr
		curr = curr.next
	}
	curr = curr.add_at_start(data)
	prev_node.next = curr
	curr.prev = prev_node
	if position == int(doubly_size){
		back = curr
	}
	return front, back
}

func doubly_new_node(data int)*doubly_node{
	doubly_size++
	new_node := new(doubly_node)
	new_node.data = data
	new_node.next = nil
	new_node.prev = nil
	return new_node
}

func (front *doubly_node)delete_at_start(back *doubly_node)(*doubly_node, *doubly_node){
	if doubly_size == 0{
		fmt.Print("List is empty\n")
		return front, back
	}
	doubly_size--
	if doubly_size == 0{
		return nil, nil
	}
	to_delete := front
	front = front.next
	to_delete.next = nil
	front.prev = nil
	return front, back
}

func (front *doubly_node)delete_at_middle(back *doubly_node, position int)(*doubly_node, *doubly_node){
	if doubly_size == 0{
		fmt.Print("List is empty\n")
		return front, back
	}else if position == 0 || position > int(doubly_size){
		fmt.Print("Invalid Position\n")
		return front, back
	}else if position == 1{
		return front.delete_at_start(back)
	}else if position == int(doubly_size){
		return front.delete_at_last(back)
	}
	doubly_size--
	to_delete := front
	for i := 1; i < position; i++{
		if to_delete.next != nil{
			to_delete = to_delete.next
		}
	}
	to_delete.prev.next = to_delete.next
	to_delete.next.prev = to_delete.prev
	return front, back
}

func (front *doubly_node)delete_at_last(back *doubly_node)(*doubly_node, *doubly_node){
	if doubly_size == 0{
		fmt.Print("List is empty\n")
		return front, back
	}
	doubly_size--
	if doubly_size == 0{
		return nil, nil
	}
	to_delete := back
	back = back.prev
	to_delete.prev = nil
	back.next = nil
	return front, back
}

func (front *doubly_node)print_list_forward(){
	fmt.Print("\nNode: ")
	for temp := front; temp != nil; temp = temp.next{
		fmt.Printf("%d -> ", temp.data)
	}
	fmt.Print("NULL\n")
}

func (back *doubly_node)print_list_backward(){
	fmt.Print("\nNode: ")
	for temp := back; temp != nil; temp = temp.prev{
		fmt.Printf("%d -> ", temp.data)
	}
	fmt.Print("NULL\n")
}


func doubly_linked_list(){
	var front, back *doubly_node = nil, nil
	for{
		fmt.Print("\nEnter your choice: ")
		var input uint
		fmt.Scanf("%d\n", &input)
		if input == 1{
			var data int
			fmt.Print("Enter the value of node: ")
			fmt.Scanf("%d\n", &data)
			if front == nil{
				front = doubly_new_node(data)
				back = front
			}else{
				node := doubly_new_node(data)
				back.next = node
				node.prev = back
				back = node
			}
		}else if input == 2{
			var data int
			fmt.Print("Enter the value of node: ")
			fmt.Scanf("%d\n", &data)
			front = front.add_at_start(data)
			if doubly_size == 1{
				back = front
			}
		}else if input == 3{
			var data, pos int
			fmt.Print("Enter the value and position of node: ")
			fmt.Scanf("%d %d\n", &data, &pos)
			front, back = front.add_at_middle(data, pos, back)
		}else if input == 4{
			front, back = front.delete_at_start(back)
		}else if input == 5{
			front, back = front.delete_at_last(back)
		}else if input == 6{
			var pos int
			fmt.Print("Enter the position of node to delete: ")
			fmt.Scanf("%d\n", &pos)
			front, back = front.delete_at_middle(back, pos)
		}else if input == 7{
			fmt.Printf("Size of list = %d\n", doubly_size)
		}else if input == 8{
			front.print_list_forward()
		}else if input == 9{
			back.print_list_backward()
		}else if input == 10{
			os.Exit(0)
		}
	}

}