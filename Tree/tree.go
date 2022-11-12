package main

import "fmt"

var count = 0

type Node struct {
	key   int
	Left  *Node
	Right *Node
}

func (n *Node) Insert(k int) {
	if n.key < k {
		// move right
		if n.Right == nil {
			n.Right = &Node{key: k}

		} else {
			n.Right.Insert(k)
		}
	} else if n.key > k {
		//move left
		if n.Left == nil {
			n.Left = &Node{key: k}
		} else {
			n.Left.Insert(k)

		}
	}
}

func (n *Node) Search(k int) bool {
	count++
	if n == nil {
		return false
	}

	if n.key < k {
		//movr right
		return n.Right.Search(k)
	} else if n.key > k {
		return n.Left.Search(k)
	}
	return true
}

func main() {

	tree := &Node{key: 100}
	tree.Insert(50)
	tree.Insert(100)
	tree.Insert(1000)
	tree.Insert(510)
	tree.Insert(10)
	tree.Insert(1000)

	fmt.Println(tree)
	fmt.Println(tree.Search(100), count)
	count = 0
	fmt.Println(tree)
	fmt.Println(tree.Search(200), count)

}
