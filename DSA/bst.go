package main

import "fmt"

type Node struct {
	value int
	right *Node
	left  *Node
}

type BST struct {
	root *Node
}

func (bst *BST) add(val int) {
	if bst.root == nil {
		bst.root = &Node{value: val, right: nil, left: nil}
		return
	}
	current := bst.root

	for current != nil {
		if val > current.value {
			if current.right == nil {
				current.right = &Node{value: val, right: nil, left: nil}
				break
			}
			current = current.right
			continue
		}
		if current.left == nil {
			current.left = &Node{value: val, right: nil, left: nil}
			break
		}
		current = current.left
	}
}

func (bst *BST) levelTraversal() {
	if bst.root == nil {
		fmt.Println("Root is nil")
		return
	}
	queue := make([]*Node, 0)
	queue = append(queue, bst.root)

	for len(queue) > 0 {
		current := queue[0]
		fmt.Println(current.value)
		queue = queue[1:]
		if current.left != nil {
			queue = append(queue, current.left)
		}

		if current.right != nil {
			queue = append(queue, current.right)
		}
	}
}

func (bst *BST) search(key int) (int, error) {
	if bst.root == nil {
		return 0, fmt.Errorf("Tree is empty!")
	}

	queue := make([]*Node, 0)
	queue = append(queue, bst.root)

	for len(queue) > 0 {
		current := queue[0]
		if current.value == key {
			return key, nil
		}
		queue = queue[1:]
		if key > current.value {
			if current.right == nil {
				return 0, fmt.Errorf("Key %d not found in tree.", key)
			}
			queue = append(queue, current.right)
			continue
		}
		if current.left == nil {
			return 0, fmt.Errorf("Key %d not found in tree.", key)
		}
		queue = append(queue, current.left)
	}
	return 0, fmt.Errorf("Key %d not found in tree.", key)
}

func inOrder(node *Node) {
	if node == nil {
		return
	}
	inOrder(node.left)
	fmt.Print(node.value, " ")
	inOrder(node.right)
}

func preOrder(node *Node) {
	if node == nil {
		return
	}
	fmt.Print(node.value, " ")
	preOrder(node.left)
	preOrder(node.right)
}

func postOrder(node *Node) {
	if node == nil {
		return
	}
	postOrder(node.left)
	postOrder(node.right)
	fmt.Print(node.value, " ")
}

func BSTDriver() {
	values := [...]int{2, 45, 67, 13, 25, 34, 10}
	bst := new(BST)

	for _, val := range values {
		bst.add(val)
	}

	// bst.levelTraversal()
	// postOrder(bst.root)
	key, err := bst.search(34)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(key)
	}
}
