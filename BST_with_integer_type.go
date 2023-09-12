package main

import (
	"fmt"
)


type Node struct { 
	Key int // Key of the node 
	// In the `Node` struct , `Left *Node` and `Right *Node` are pointer fields . 
	//They hold the memory address of another `Node` struct, allowing to created a linked structure, in our case 
	// a BST. 
	Left *Node // Pointer to the left child . 
	Right *Node // Pointer tot the right child .
}


func (n *Node) Insert (k int) { 
	if n.Key < k { // If current node 's key is less than k 
		if n.Right == nil {  // if right child is nil . 
			n.Right = &Node{Key: k}  // Create a new  right child . 
		} else { 
			n.Right.Insert(k) // Otherwise, insert into the right subtree
  	}
	} else if n.Key > k {    // If current node' s key is greater than k
	if n.Left == nil {       // If left child is nil 
		n.Left = &Node{Key:k}  // Create  a new left child .  
	} else { 
		n.Left.Insert(k) // Otherwise , insert in the left subtree . 
	}
 }
}

// The `Search` function looks for a node with the given key and returns true if found. 
func (n *Node) Search (k int) bool  { 
	if n == nil {  // If the node is nil 
		return false 
	} 
	if n.Key == k {  // If the key matches . 
	return true 
} else if n.Key < k {  // If the key is less than k .
	return  n.Right.Search(k)  // Search in the right subtree . 
} else { 
	return n.Left.Search(k) // Search in the left subtree . 
 }
}


func main() {
	 tree := &Node{Key: 100} // Initialmize the tree with root key 100 
	tree.Insert(200) // Insert 200 in the tree . 
	tree.Insert(50)  // Insert 50 in the tree . 
	tree.Insert(202) // Insert 202 in the tree .
	tree.Insert(999)  // Insert 8 in the tree . 

	fmt.Println("Is the number 200 in the tree ?" , tree.Search(200)) 
	fmt.Println("Is the number 7 in the tree ?", tree.Search(7)) 
}
