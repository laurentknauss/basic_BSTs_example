
package main

import (
	"testing"
)

func TestInsertAndSearch(t *testing.T) { 
	tree := &Node{Key:100} // initializes a new binary search tree with a root node that  has a value of '100' . 
	tree.Insert(201) // inserts a new node with the key value of `200` into the tree . 
	tree.Insert(59) // inserts a new node with a key value of `50` into the tree . 
	

	if !tree.Search(100) {  // This line checks if the `Search` method can find a node witht the key value of `100`. 
	// If it can not, the test will not pass . 
		t.Errorf("Expected to find 100 in the tree")
	}
	if !tree.Search(201) {   // This line checks if the `Search` method can find a node with a key value of `201`. 
	// If it can not, the test will not pass . 
		t.Errorf("Expected to find 201 in the tree")
	}
	if !tree.Search(59) {  // This line checks if the `Search` method  can find a node witht he key value of `59` . 
	// If it can not , the test will not pass . 
		t.Errorf("Expected to find 59 in the tree")
	}
	if tree.Search(999) {  // This line checks if the `Search` method errorneously finds a node with the key value of `999` which was not inserted in the tree . 
	// If a node with key value of `999` is found , which would be incorrect, an error message is logged. Verifying that the method correctly handles 
	// absent keys is a good practice . 
		t.Errorf("Did not expect to find 999 in the tree")
	}
}
 