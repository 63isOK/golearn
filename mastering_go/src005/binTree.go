package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Tree ...
// 节点
type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// 前序遍历
func traverse(t *Tree) {
	if t == nil {
		return
	}

	traverse(t.Left)
	fmt.Print(t.Value, " ")
	traverse(t.Right)
}

func create(n int) *Tree {
	var t *Tree
	rand.Seed(time.Now().Unix())
	for i := 0; i < 2*n; i++ {
		tmp := rand.Intn(n * 2)
		t = insert(t, tmp)
		fmt.Print(tmp, " ")
	}

	return t
}

func insert(t *Tree, v int) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}

	if v == t.Value {
		return t
	}

	if v < t.Value {
		t.Left = insert(t.Left, v)

		return t
	}

	t.Right = insert(t.Right, v)

	return t
}

func main() {
	t := create(5)
	fmt.Println()
	fmt.Println("root:", t.Value)
	traverse(t)
	fmt.Println()
	// t = insert(t, -10)
	// t = insert(t, -2)
	// traverse(t)
	// fmt.Println()
	fmt.Println("root:", t.Value)
}
