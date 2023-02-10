package main

import (
	"fmt"
	"strings"
)

type Node struct {
	left *Node
	right *Node
	val int
}

func print_bst(root *Node, tag string) {
	if root != nil {
		fmt.Print(tag, root.val, "  ")
		print_bst(root.left, "left")
		print_bst(root.right, "right")
	}
}

func input(btArray []int) *Node {
	if len(btArray) == 0 {
		return nil
	}

	val := btArray[0]
	rightBegin := -1

	root := Node { val: val }

	for i:=1 ; i<len(btArray) ; i++ {
		if btArray[i] > val {
			rightBegin = i
			break
		}
	}

	if rightBegin > 1 {
		left := input(btArray[1 : rightBegin])
		root.left = left
	}

	if rightBegin > -1 {
		right := input(btArray[rightBegin : ])
		root.right = right
	}

	return &root
}

func traverse_levels(root *Node, levels *[][]int, current_level int) {
	if current_level >= len(*levels) {
		row := []int{}
		*levels = append(*levels, row)
	}

	(*levels)[current_level] = append((*levels)[current_level], root.val)

	if root.left != nil {
		traverse_levels(root.left, levels, current_level + 1)
	}

	if root.right != nil {
		traverse_levels(root.right, levels, current_level + 1)
	}
}

func main() {
	b := []int { 10, 5, 2, 7, 20, 8, 22 }
	bt := input(b)
	levels := [][]int {}
	print_bst(bt, "root")

	traverse_levels(bt, &levels, 0)

    for level_index, row := range levels {
		fmt.Println()
		fmt.Print(level_index)
		fmt.Println()
		fmt.Print(strings.Repeat(" " , len(b) * 3  - level_index * 3))
		for _ , node := range row  {
			fmt.Print(node, "  ")
		}
		fmt.Println("")
    }
}
