package sol

import "fmt"

type TreeNode struct {
	Val int
	L   *TreeNode
	R   *TreeNode
}
type Node struct {
	Val  int
	Next *Node
	Prev *Node
}

func BFT(root *TreeNode, k int) int {
	// create a queue
	queue := append([]*TreeNode{}, root)
	rootQueue := append([]*TreeNode{})
	refMap := make(map[int]*Node)
	head := &Node{Val: root.Val}
	refMap[root.Val] = head
	count := 0
	for len(queue) != 0 {
		visit := queue[0]
		visitRoot := &TreeNode{} // root value
		if len(rootQueue) > 0 {
			visitRoot = rootQueue[0]
			rootQueue = append(rootQueue[1:])
		}
		count++
		queue = append(queue[1:])
		if count != 1 {
			// fmt.Println(visitRoot.Val, visit.Val)
			visitNode := &Node{Val: visit.Val}
			refMap[visit.Val] = visitNode
			if visitRoot.Val > visit.Val {
				if v, ok := refMap[visitRoot.Val]; ok {
					// v.Next = visitNode.Next
					visitNode.Next = v
					if head.Val == visitRoot.Val {
						head = visitNode
					} else {

						if v.Prev != nil {
							v.Prev.Next = visitNode
						} else {
							v.Prev = visitNode
						}
					}
				}
			} else {
				if v, ok := refMap[visitRoot.Val]; ok {
					visitNode.Next = v.Next
					v.Next = visitNode
					visitNode.Prev = v
				}
			}
		}
		if visit.L != nil {
			rootQueue = append(rootQueue, visit)
			queue = append(queue, visit.L)
		}
		if visit.R != nil {
			rootQueue = append(rootQueue, visit)
			queue = append(queue, visit.R)
		}
	}
	return FindMinK(head, k)
}
func FindMinK(head *Node, k int) int {
	result := 0

	for count := 0; count < k-1; count++ {
		head = head.Next
	}
	result = head.Val
	return result
}
func Transerval(head *Node) {
	for ; head != nil; head = head.Next {
		fmt.Printf("%d ->", head.Val)
	}
	fmt.Println("nil")
}
func Solution(root TreeNode, k int) int {
	result := 0
	result = BFT(&root, k)
	return result
}
