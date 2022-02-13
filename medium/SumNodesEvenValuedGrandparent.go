package medium

/*
	Given the root of a binary tree, return the sum of values of nodes with an even-valued grandparent. If there are no nodes with an even-valued grandparent, return 0.
 	A grandparent of a node is the parent of its parent if it exists.

	 Input: root = [6,7,8,2,7,1,3,9,null,1,4,null,null,null,5]

	Output: 18
	Explanation: The red nodes are the nodes with even-value grandparent while the blue nodes are the even-value grandparents.

*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumEvenGrandparent(root *TreeNode) int {
	return dfs(root, nil, nil)

}

func dfs(root, parent, grandParent *TreeNode) (res int) {
	if root == nil {
		return 0
	}
	if grandParent != nil && grandParent.Val%2 == 0 {
		res += root.Val
	}

	return res + dfs(root.Left, root, parent) + dfs(root.Right, root, parent)
}
