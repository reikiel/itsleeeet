/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    if root == nil {return true}

    return depth(root) != -1

}

func depth(root *TreeNode) int {
    if root == nil { return 0 }

    left, right := depth(root.Left), depth(root.Right)

    if (absInt(left-right) > 1 || left == -1 || right == -1) {
        return -1
    }

    return maxInt(left, right) + 1
}

func absInt(a int) int {
    if a < 0 {
        return -a
    }
    return a
}

func maxInt(a, b int) int {
    if a < b {
        return b
    }
    return a
}