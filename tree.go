package main

type Tree struct {
	node *Node
}

type Node struct {
	value int
	left  *Node
	right *Node
}

func (t *Tree) insert(value int) {
	if t.node == nil {
		t.node = &Node{value: value}
	} else {
		t.node.insert(value)
	}
}

func (n *Node) insert(value int) {
	if value <= n.value {
		if n.left == nil {
			n.left = &Node{value: value}
		} else {
			n.left.insert(value)
		}
	} else {
		if n.right == nil {
			n.right = &Node{value: value}
		} else {
			n.right.insert(value)
		}
	}
}

func (t *Tree) Display() {
	if t.node == nil {
		println("Tree is empty")
	} else {
		for t.node != nil {
			println(t.node.value)
			t.node = t.node.left
		}
	}
}
