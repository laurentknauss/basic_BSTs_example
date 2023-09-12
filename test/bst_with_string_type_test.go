
package main

import (
	"testing"
)

func TestInsertAndSearch(t *Testing.T) { 
	tree := &Node{Key:100}
	tree.Insert(200)
	tree.Insert(50)

	if !tree.Search(100) { 
		t.Errorf("Expected to find 100 in the tree")
	}
	if !tree.Search(200) { 
		t.Errorf("Expected to find 200 in the tree")
	}
	if !tree.Search(50) { 
		t.Errorf("Expected to find 50 in the tree")
	}
	if tree.Search(999) { 
		t.Errorf("Did not expect to find 999 in the tree")
	}
}
