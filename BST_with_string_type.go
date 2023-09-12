package main

import (
	"fmt"
)


type Node struct { 
	Key string 
	Left *Node
	Right *Node
}

// Insert adds a new node witht he given key to the tree . 
func (n *Node) Insert(k string) { 
	if n.Key < k {  // String comparison 
	if n.Right == nil {
		n.Right = &Node{Key: k}
	} else {
		n.Right.Insert(k)
	}
} else if n.Key > k {
	if n.Left == nil {
		n.Left = &Node{Key: k}
	} else {
		n.Left.Insert(k)
	}
}
}
// Search looks for a node with a given key and returns `true` if found . 
func (n *Node) Search ( k string) bool { 
	if  n == nil {  // 
		return false 
} 
if n.Key == k { 
	return true 
} else if n.Key < k { 
	return n.Right.Search(k) 
} else { 
	return n.Left.Search(k) 
}

}
func main() {
	tree := &Node{Key:"apple"} 
	tree.Insert("banana") 
	tree.Insert("cherry") 
	tree.Insert("kiwi") 
	tree.Insert("strawberry") 

	fmt.Println("Is 'orange' in the tree ?", tree.Search("orange")) 
	fmt.Println("Is 'kiwi' in the tree ?", tree.Search("kiwi")) 
}
