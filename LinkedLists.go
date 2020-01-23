package main

import (
	"fmt"
)

type Node struct {
	Value int
	Next *Node
}

type List struct {
	Head *Node
}

func ( l *List) Append(n Node) {
	cur := l.Head
	for cur.Next != nil {
	cur = cur.Next
	}
	cur.Next = &n
}

func (l *List) PrintList(){
	i := l.Head
	fmt.Println("(Head)")
	for i != nil {
	fmt.Println(i.Value, " -> ")
	i = i.Next 
	}
	fmt.Println("Nil")
}

func (l *List) Reverse() {
	i := l.Head
	var prev *Node = nil
	for i != nil {
		next := i.Next
		i.Next = prev
		prev = i
		i = next
	}
	l.Head = prev
}

func main() {
	list := List{&Node{1, nil}}
	list.Append(Node{2,nil})
	
	list.PrintList()
	list.Reverse()
	list.PrintList()
}
