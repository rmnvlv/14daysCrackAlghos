package day08

// Definition for a Node.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

// Runtime: 7 ms, Memory Usage: 6.3 MB
func connect(root *Node) *Node {

	if root == nil {
		return root
	}

	dummy := root

	for dummy.Left != nil {
		field := dummy

		for field != nil {
			field.Left.Next = field.Right

			if field.Next != nil {
				field.Right.Next = field.Next.Left
			}

			field = field.Next
		}

		dummy = dummy.Left
	}

	return root
}
