/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    if root == nil {
        return true
    }

    return depth(root) != -1
}

func depth(root *TreeNode) int {
    if root == nil { return 0 }
    left := depth(root.Left)
    right := depth(root.Right)

    if left == -1 || right == -1 || abs(left-right) > 1 {
        return -1
    }

    return max(left, right)+1

}

func max(a,b int) int {
    if a > b {
        return a
    }
    return b
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}