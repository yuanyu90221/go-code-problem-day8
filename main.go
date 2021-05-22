package main

import (
	"fmt"

	"github.com/yuanyu90221/go-code-problem-day8/sol"
)

func main() {
	bst := &sol.TreeNode{
		Val: 5,
		L: &sol.TreeNode{
			Val: 3,
			L: &sol.TreeNode{
				Val: -1,
			},
		},
		R: &sol.TreeNode{
			Val: 9,
			L: &sol.TreeNode{
				Val: 6,
				R: &sol.TreeNode{
					Val: 7,
				},
			},
			R: &sol.TreeNode{
				Val: 10,
			},
		},
	}
	result := sol.BFT(bst, 4)
	fmt.Println(result)
}
