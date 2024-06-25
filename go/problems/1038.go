// 1038 - Binary Search Tree to Greater Sum Tree
package problems

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rvrsInOrder(node *TreeNode, carry *int) {
	if node == nil {
		return
	}
	// inorder reverse - right tree first
	rvrsInOrder(node.Right, carry)

	*carry += node.Val
	node.Val = *carry

	rvrsInOrder(node.Left, carry)
}

func BstToGst(root *TreeNode) *TreeNode {
	var carry int
	rvrsInOrder(root, &carry)
	return root

}
