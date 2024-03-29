package main

import (
	"errors"
	"fmt"
	"strings"
)

type BinarySearchTree struct {
	value int
	left  *BinarySearchTree
	right *BinarySearchTree
}

func New(value int) *BinarySearchTree {
	tree := &BinarySearchTree{}
	tree.value = value

	return tree
}

func (tree *BinarySearchTree) Display(level int) {
	if tree == nil {
		return
	}
	tree.right.Display(level + 1)
	fmt.Println(strings.Repeat("   ", level), tree.value)
	tree.left.Display(level + 1)
}

func (tree *BinarySearchTree) Insert(value int) *BinarySearchTree {
	if tree == nil {
		return &BinarySearchTree{value: value}
	}

	if value > tree.value {
		tree.right = tree.right.Insert(value)
	} else if value < tree.value {
		tree.left = tree.left.Insert(value)
	}

	return tree
}

func (tree *BinarySearchTree) Search(value int) (result int, err error) {
	if tree == nil {
		return 0, errors.New("cannot find value")
	}

	if tree.value == value {
		return tree.value, nil
	}

	if value > tree.value {
		return tree.right.Search(value)
	}

	if value < tree.value {
		return tree.left.Search(value)
	}

	return 0, errors.New("cannot find value")
}

func (tree *BinarySearchTree) Delete(value int) *BinarySearchTree {
	if tree == nil {
		return nil
	}

	if value == tree.value {
		if tree.right == nil && tree.left == nil {
			return nil
		}

		if tree.right != nil && tree.left == nil {
			return tree.right
		}

		if tree.right == nil && tree.left != nil {
			return tree.left
		}

		if tree.right != nil && tree.left != nil {
			minNode := tree.FindMinNode()
			newTree := tree.Delete(minNode.value)

			newTree.value = minNode.value

			return newTree
		}
	}

	if value > tree.value {
		tree.right = tree.right.Delete(value)
	}

	if value < tree.value {
		tree.left = tree.left.Delete(value)
	}

	return tree
}

func (tree *BinarySearchTree) FindMinNode() *BinarySearchTree {
	currentNode := tree

	for currentNode.left != nil {
		currentNode = currentNode.left
	}

	return currentNode
}

func (tree *BinarySearchTree) Reverse() *BinarySearchTree {
	if tree == nil {
		return nil
	}

	temp := &BinarySearchTree{value: tree.value}
	temp.right = tree.left.Reverse()
	temp.left = tree.right.Reverse()

	return temp
}

func main() {
	tree := New(0)
	tree = tree.Insert(2)
	tree = tree.Insert(1)
	tree = tree.Insert(3)
	tree = tree.Insert(-2)
	tree = tree.Insert(-1)
	tree = tree.Insert(-3)
	tree = tree.Delete(1)
	tree = tree.Delete(-2)
	result, err := tree.Search(0)

	if err != nil {
		fmt.Printf("Cannot find value")
		return
	}

	fmt.Println(result)
	fmt.Println("------------")
	tree.Display(0)
}
